package keeper

import (
	"fmt"
	"github.com/btcsuite/btcd/wire"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/babylonchain/babylon/x/btclightclient/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// InsertHeader inserts a btcd header into the header state
func (k Keeper) InsertHeader(ctx sdk.Context, header *wire.BlockHeader) error {
	headerHash := header.BlockHash()
	headerExists := k.HeadersState(ctx).HeaderExists(&headerHash)
	if headerExists {
		return types.ErrDuplicateHeader.Wrap("header with provided hash already exists")
	}

	parentExists := k.HeadersState(ctx).HeaderExists(&header.PrevBlock)
	if !parentExists {
		return types.ErrHeaderParentDoesNotExist.Wrap("parent for provided hash is not maintained")
	}

	height, err := k.HeadersState(ctx).GetHeaderHeight(&header.PrevBlock)
	if err != nil {
		// Height should always exist if the previous checks have passed
		panic("Height for parent is not maintained")
	}

	parentWork, err := k.HeadersState(ctx).GetHeaderWork(&header.PrevBlock)
	if err != nil {
		// Work should always exist if the previous checks have passed
		panic("Work for parent is not maintained")
	}

	headerWork := types.CalcWork(header)
	cumulativeWork := types.CumulativeWork(headerWork, parentWork)

	// Create the header
	k.HeadersState(ctx).CreateHeader(header, height+1, cumulativeWork)
	return nil
}
