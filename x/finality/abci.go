package finality

import (
	"context"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/babylonchain/babylon/x/finality/keeper"
	"github.com/babylonchain/babylon/x/finality/types"
)

func BeginBlocker(ctx context.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	return nil
}

func EndBlocker(ctx context.Context, k keeper.Keeper) ([]abci.ValidatorUpdate, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// if the BTC staking protocol is activated, i.e., there exists a height where a finality provider
	// has voting power, start indexing and tallying blocks
	if _, err := k.BTCStakingKeeper.GetBTCStakingActivatedHeight(ctx); err == nil {
		// index the current block
		k.IndexBlock(ctx)
		// tally all non-finalised blocks
		k.TallyBlocks(ctx)

		// detect inactive finality providers if there are any
		// height for examining is determined by the current height - params.LivenessDelay
		// this is we allow finality signatures to be received after quorum
		heightToExamine := sdk.UnwrapSDKContext(ctx).HeaderInfo().Height - k.GetParams(ctx).LivenessDelay
		if heightToExamine >= 1 {
			k.HandleLiveness(ctx, heightToExamine)
		}
	}

	return []abci.ValidatorUpdate{}, nil
}
