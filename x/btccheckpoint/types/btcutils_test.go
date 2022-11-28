package types_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	bbn "github.com/babylonchain/babylon/types"
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btcchaincfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

//nolint:unused
func hashFromString(s string) *chainhash.Hash {
	hash, e := chainhash.NewHashFromStr(s)
	if e != nil {
		panic("Invalid hex sting")
	}

	return hash
}

// Sanity test checking mostly btcd code, that we can realy parse bitcoin transaction
func TestBtcTransactionParsing(t *testing.T) {
	// Few randomly chosed btc valid btc transactions
	tests := []struct {
		bitcoinTransactionHex string
		txHash                string
	}{
		{"0100000001bff216373dcb167b61faea34694713de29115b69690b9f9b6f6019c5e6edc7eb010000006b483045022100d68de754b884bb3b99905f57e856f8111152cdc597c469c83c8340f0e7a30044022019335a825a2a5ddea9e45e28b3eaeb3a042d9377b2a9700f4ddd2262d77174e5012103c164d4864ae50334ea2bdc9d7aaad0418574c96062d5b739b505cd985fc3bd90ffffffff02404b4c000000000017a914df6122d5edb7697b466f7201a862b4fd67ea2ac1871055aa13000000001976a914ad595785759e25fa948505b52f1351bdbed1bba288ac00000000", "c013cd25a9e73b678eb8c8a7304890beb7b29dd18864f0379a562335d3c37a8b"},
		{"0100000001f0606dd02119ee3660156250762a26bf17be8edef3789dcf67a12586a8be4676000000006a47304402202e02ef95a2958bd302668a0b64e7ed8e886efaea7103b077f1882c576ae443db022059f8d853d17b5bdeedbe0cd41be98ad412237743c8bf07dbb6e90d4627a02ac1012103e13e1a05703b972b47e228bc9459be59e09dad0bc3a54d7d8d9dc20f63b4af48ffffffff02404b4c000000000017a914df6122d5edb7697b466f7201a862b4fd67ea2ac18750ce9000000000001976a914c5b58a131b5a2d751a8176a8f54d1f039ebf9d6188ac00000000", "fb2a3934e40d82744150fd3db3b3b5c09ad5f33f2b873a1b058da0acb5a3838d"},
		{"01000000027c918745a69ce4ff8e674fbf059fea3dfdf222f2c15df217b9c4bb218ddb0d38010000008b4830450221008259debd856a0afaf33c918351a75759742061738c5fc0790b5813cbef95566a022022ff0311284290f3a84b4b1d6fd550ed5f71c9d522551e2f225537138050c9a40141048aa0d470b7a9328889c84ef0291ed30346986e22558e80c3ae06199391eae21308a00cdcfb34febc0ea9c80dfd16b01f26c7ec67593cb8ab474aca8fa1d7029dffffffff83a1289fe5b1c5b4bcbd5beeac029fc88ff87272b8deb431155240b952bc00b8000000008a473044022071a59dfccc514273fd9a3f765ea58c7a24cd81a34c8158b636e3507ee374531702200fc40b8689b43469c1607da8bb78f3565e98ad87c7196f00a497dcd17dea57720141048aa0d470b7a9328889c84ef0291ed30346986e22558e80c3ae06199391eae21308a00cdcfb34febc0ea9c80dfd16b01f26c7ec67593cb8ab474aca8fa1d7029dffffffff01287d0200000000001976a91436a5ee46338acf885538ebd709a810b361c93a4388ac00000000", "0418fb71c9c7fc1fb1d383149dc0c01bdc6894d867535b2758ceae57e83ab805"},
		{"0100000001ab8ead1a89e7959d1aa325b1b619ccb164bc9b77cc6d2decfb064f120951af5b000000006a47304402200ad6ba00c17d78095e1e46203e438d02179345b87a4f9af2eed985ddf7b7432a0220738e94e64ba6a576166cec83dbcdc7e827c1479eab1df50f8cf323fa71db4b89012103bb6e200caa87be8296d5b4398684729f27c7d44f9b29ab96640793fdd23a864affffffff1437020000000000001976a914425245544852454e0a0a6168696d73612d64657688ac37020000000000001976a91412cd026060604265666f72652049206c6566742088ac37020000000000001976a914536f757468204166726963612c2061206c616e6488ac37020000000000001976a9142049206c6f76652070617373696f6e6174656c7988ac37020000000000001976a9142c2077652068616420616e20656d657267656e6388ac37020000000000001976a91479206d656574696e67206f66207468652045786588ac37020000000000001976a91463757469766520436f6d6d6974746565206f662088ac37020000000000001976a91474686520536f757468204166726963616e20436f88ac37020000000000001976a914756e63696c206f6620436875726368657320776988ac37020000000000001976a914746820746865206c656164657273206f66206f7588ac37020000000000001976a91472206d656d6265722063687572636865732e205788ac37020000000000001976a914652063616c6c656420746865206d656574696e6788ac37020000000000001976a9142062656361757365206f6620746865206465657088ac37020000000000001976a914656e696e672063726973697320696e206f75722088ac37020000000000001976a9146c616e642c2077686963682068617320636c616988ac37020000000000001976a9146d6564206e6561726c7920323030206c6976657388ac37020000000000001976a9142074686973207965617220616c6f6e652e60606088ac37020000000000001976a9140a0a2d2d4465736d6f6e64205475747518e88ffa88ac37020000000000001976a914ac0500000000000000000000000000000000000088ac237c4400000000001976a9148fe22047956d1510f84018a3f3ab91686b82f2ea88ac00000000", "cac99ea9fc912ef42b4059add3d853c722d32be8b903fef24e7fd08f05960eab"},
		//this one has op_return data
		{"01000000018dc21dfe0dae9c213b015c496da276fd090ad4ba22386f53b28985f374e50b28020000006a47304402207ac812172733436bd0ee318aa3064da7e49e212f59a3e62f7a6f3b6f11220b5f0220690b7174a607e64eed5e2ae3622f5d5fcffc7f947523d23bd428ade7897f55220121030229d577f4d7b45d50b55dd93e5024561baa6ddf3ba29188d43b9b357618e896ffffffff03c8190000000000001976a914c4a0be51e3280054bcfc7471182ddbaf3bb3c58e88ac0000000000000000396a374f4101000101303030303030303030303030303031333130326574597668696c646c6d4762746a54694c3648494f776a4b64756e6f4d559082b302000000001976a914c4a0be51e3280054bcfc7471182ddbaf3bb3c58e88ac00000000", "af5575e0a1ae1ca1bbca55f74f7504ef9bc50280a7da9bd5ce210d8a4392ea33"},
		// op_return and 4 outputs
		{"0100000001b9e93d75c8bb80716709ff41d6a89585c9408b50cf4f5c6f6404a7e8f23d3b7006000000fb00473044022036f79054547c3696b87abc30bbfd4815f49199c8d3adbacfe798ce2c91775264022036dfd5a59ef722c568580f818cee19cf4da9d9be638833573fd811d7888d9291014630430220239b77e69f536f5f9e22bf6286d3f7bddb9493e087583563e090dfaeab1dbe29021f2dd263ca25643469652bba0cf1cc2c753318acc4599e275d163408943a3d07014c69522102081bc3fe67cf4fe2397c9c232993aaff3ddbad39982253d3e09ef8f6655d958721037d8ab0e3bc36818067d042b9a714363ee36a6e4e3c5c3c35b785377a43cef0d22103b1a264f3fd81af6bd775aaae218444302945c78a711e05679107ffc713be12fc53aeffffffff0500000000000000000b6a095349474e4154555245aee353000000000017a914715b26c59e87d9247fb0b0171e85c32f5f0a6f7787fb750e000000000017a914ff718c1a44a41fe1decc73d601cb3b5f0c189377871b5b03000000000017a9144a456da09638ebb610364742f51713d379908257877c3613000000000017a914887fee10bad2a610796eefec5aeddcb607045a618700000000", "4ad274333d08cd55981b217017fa209d3c435364c94a42f666c83a551d67961b"},
	}

	for i, test := range tests {
		txBytes, _ := hex.DecodeString(test.bitcoinTransactionHex)
		hashBytes, _ := chainhash.NewHashFromStr(test.txHash)

		tx, e := btcctypes.ParseTransaction(txBytes)

		if e != nil {
			t.Errorf("Failed to parse valid bitcoin transaction %d", i)
		}

		if !tx.Hash().IsEqual(hashBytes) {
			t.Errorf("Hash of decoded transaction does not match provided hash %d", i)
		}

	}
}

// All data here is taken from btc testnet
func TestParsingCorrectBtcProofs(t *testing.T) {
	tests := []struct {
		header                 string
		transactions           []string
		opReturnTransactionIdx int
		expectedOpReturnData   string
	}{
		{
			"000000200761788fd78b840d06e9b0e2e46ee8645ebf294136abe64b53430300000000006250ecc747676b963379d201e244c074f7673e17acc054a410fbdd2e9662debb08547958fcff031b705d182c",
			[]string{
				"01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2203c27a10150e6d696e65642062792062636f696e58795408100000000000cf110000ffffffff0214d917000000000017a9146859969825bb2787f803a3d0eeb632998ce4f50187bcf238090000000017a9146859969825bb2787f803a3d0eeb632998ce4f5018700000000",
				"01000000017ba216f5893308e3a6706d6a644da0e58b765cf4fe8d6dd78d14b7b771f6892601000000fdfd0000473044022075025b2858500747b3a9005b616625f0143610c06d36aee3153c52f377e82f1d02202215d5feff5daee05b81be934f38c0b01b0194874a1b4372d211a141bc46c7440148304502210083fc1e3553d9ea849387ee704fde23d97909b54e4b7b59316c702b72177775610220429c2128bb1602df56e6fdb463be6dd40acaeeebfe3cd80abb85fe312fa3185c014c69522102b4905cfaa07073953235d02a9046c76ce398a4e7d0b41b2004089698304b11cd2102eeee7cdafe05cc4cf526355e866cf8bc146a50172cf3ab48661eed00bd423a9a21038e282cca3b7851d02fe10c4c6b99a175662b3ad796daac5a9356002aab652d5853aeffffffff0600000000000000000b6a095349474e4154555245a94c87000000000017a9140c7497f8303b61239314a4c7022d5e9081130f2b877a5925000000000017a9141850de01509f098efae2b8a3cde05d6ae34767fb87759707000000000017a9143297b2f17781542470eb6143c2b70f86697b152687c24c00000000000017a9142c5ee431f022ff2ccafee158b6b9db046f5554a087c90b1e000000000017a914081e2f578f7fd6c641578945f1e432ba0de60a038700000000",
			},
			1,
			"5349474e4154555245",
		},
		{
			"0000e0208c3b3ed3aa778eaecdcbe91dae57197ce1baa0d7c33e86d00d0100000000000079ffca6c6b36348c306234dee2fe47bafd76df7e70c95cbdff3efeb81e5abe71ea88b860fcff031a45722027",
			[]string{
				"010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff240381841e0c2074657374206d696e65722012097573200909200902825401fa4184010000ffffffff02efbc97000000000017a9140d2eb00a31486c91e3dbefa13ac714e236390dad870000000000000000266a24aa21a9ed02ee31a4ff032e606a5bc1af454ddca6695a1261a69d4ddb24d6dd10cb6d3fcd0120000000000000000000000000000000000000000000000000000000000000000000000000",
				"02000000000101761d35946ece6a53b79380119ccda626a8efd5caee724d71d81f404f5a33003f00000000171600145074935eaaf3cc1f04acc64c2c4f88737ff17896feffffff02f0f20f000000000017a9147e96bcc24d343e35f857f593eca765ccdf200b17872172b38a0000000017a914647dbca76f3d3426a564361c9539aae810752e4d870247304402204e80eb98037ec577a88a4712e6b3ea81eb23541052ecd595211edce78034fbd8022048bc0507ca70261810906c818e761d6a53cb69ad11c2c19ac554047622e88f85012103521bdcb10ea983094184c8b8dc49698541d63c45d91d90771b75434d568c80fb7f841e00",
				"02000000000101b8ff4b6851eee40c958c58916604c15533e7fb4a64a9ec509a55d10167b64b290000000000feffffff023b1ccdd700000000160014484f0aa3ea9b74beb476be62936df51c2fc59b99e84610000000000016001448717afa6934ad4da8e5e827ca061d11e41f78e602473044022079d02fd7cca6aa2f3e8860987b3fae269677ea97bc1c79553d1e53b6d5329454022040ecf6545258d0702713f46d4bafe1f1eb05c1dec983b84982d1cf2c807195dd0121034120b65994ee9788e450312b312d1fcef975c0f34dadd0fbe8ed01c9c61633a880841e00",
				"0100000001e7742bc7ac999bd4b0832809534a5965cf5b53abdc75e3f840702be92b1f82dd030000006a47304402202fe1e5defb67549a2f2a8b7d754a80866583f348e2aae9a61ece13b0842b16ca0220311eacc4c8bf859ba2ad76b7093caa366cff4b4ec66fb6e14708915cad3a04e40121037435c194e9b01b3d7f7a2802d6684a3af68d05bbf4ec8f17021980d777691f1dfdffffff040000000000000000536a4c5058365b8ca81abb972accd336d8f05be0c843b5486fb72528877a5f50e898b14d3b39e011a46ff8b0ad5a933ab2caee547240cb41e7da3cf69d5f364f4ba76e06796dc5001e847e0005001e818e00010010270000000000001976a914000000000000000000000000000000000000000088ac10270000000000001976a914000000000000000000000000000000000000000088acc1e09a04000000001976a914ba27f99e007c7f605a8305e318c1abde3cd220ac88ac00000000",
				"0100000001d6541ff19b573a4742925a56552e608fc827a149e412e30018e64fc39c6ebcbb010000008b483045022100c9a4d2502164a78caefd4d3b1deac72bdf418636e2cf16bea7a051438bec8499022070dcebad8a71a1acc6fea62512ae88a9e7901d31cbeca4a1d1af10640eba538c0141045a5ddc925295b71bafbe56bf4c10e1c1bc7c3a2bf5116b72f5dd202bccc032955afc5191f626284508072d397fd0fde700ae6feb2a35c1c391b12971960e6df6ffffffff03ee050000000000001600140fb58dc4fc27d579fd59cd18d3b44f8b5df1b47b2b2be30b000000001976a914d9ea351605b36fc3a967d790132230eb7eced36688ac0000000000000000256a2302000fa26dbf437f2811124e8395d532f969f2ee83a6d0542e2f5798ce37f267f2fdaa00000000",
				"01000000000101614b0cdbc00644c6e8bb016ab669acafacab0279c4df379440cf225ed1fd2c8a0100000000ffffffff02407e05000000000017a914dc75fc89f54f9618ee4fb5ef538c3baa46adf7458719480800000000001600146bee0d3f361510aae3e9d5014f8f91db21342506024830450221009315bcf6106e8666ea4dea43f09a81550f4b481b67e20375a3acb2678de500e402201d8a79065987b1cad8d032d4e58652d6258d4e3f53b7d15b2a78029c77f595100121035ad29a75b4561eb39c4f81a6f860d032385ef56717f3b896322c9a611cb40ed300000000",
				"02000000000101bfb1f07e1055a1ec06c045322ce1d8354d59f10fea0ed3294bd12954150f404f0100000000fdffffff02142b020000000000160014f2bb4458b593dfd6d96d1474469e240121ccfc0540420f0000000000160014206f30d19f2d39dfeda1933e07a62e38ce380da10247304402202576f7adcca78e963366c7b16c3732f3401ccc42e1726ec1e675cecd1928591a022028b60a4ab7de85112252aacaca74f3caf4927a2932c719cecb008194a51a45ad012102b65c0805712db5158e1554297fe51e468e8bbb25464d3fed183d8e7b969b539e7d841e00",
				"020000000001032c5faf7fc3751f79f93ed389db1af6394fffede942d51fcec61d15bb5829302d01000000171600147af2e87a43f5370b17bb0bd0ea18849e123ba61afeffffffad6d0fe044f3440347fa07e2fe07fac2d1a585f0a66b9a23265d968667aa21ec0000000017160014b5d5a81ba37d47ad5fe276b23f3973f56b70a459feffffff2b564e1dd88cbddbd79ba2b5467f64cbbec439c008738906e361f1b2f5eb0fe40000000000feffffff02801a060000000000160014407d31e030df0f64a73f142a0feecd4875a0626c32810f00000000001600149cffa54d75caa5cb3897efde3ef9017d593e203a0247304402200bf2fcd9dec73082763bef3999f531f9b6b5f707535596244623f874973de610022013ac74dd9a5b6eb7300e35bea44546f62b726ec4b277342920d2378df34361a501210253d9126db8349ae5550fadd02ebe612fea2dd239ba32f1029342d407d4ab6fba02473044022048b1d568ee52de84ca71dddab0503e1e8f4829a6478a1d69baeab8896ad5c9230220757e3255cf415e067f843396e5226ff6ed8facb45013773e18aa9bb0243791f50121038a51c7db5ab7377e8d5f9a06be6b6b46c900df51af9495a79f3ef01a46c71f4902473044022005030ff6d9539774eb1a03eb3ae06e73c15011404076cac08dc2fbd4d54b67a0022052d902113c9ed40c326bea976d3f0ce475fc3a575ea55e9bf158d410fc842f46012103c2f8ebfebd41bfa467ab63f3ef7e734c64ee38a46fc2f4207f3aca935bd170be7f841e00",
				"010000000001010f90038c463b5f6fc0593ec8110951c665052b2631e5c0978f65e48706d195b30100000000ffffffff02b5264408000000001600147fa2f0e860094b79fda8ac58245e2c9c5c1fbb1f8c5a0000000000001976a9141d7cb15cc4393354b4007e97161c58c1e16f0c0d88ac02473044022030aa6f4718e6dc75e0448b6de4c4d8a2a0c1c49be5c5eca1d3902b1b8a44c9da02201982eead66664a27637b46b90fc88bd80126208381700be6be327f0c8c79386d0121029459153248d701d7402b0ef3dddfde42f3867a52f172fa383746287014dd4a2e00000000",
			},
			3,
			"58365b8ca81abb972accd336d8f05be0c843b5486fb72528877a5f50e898b14d3b39e011a46ff8b0ad5a933ab2caee547240cb41e7da3cf69d5f364f4ba76e06796dc5001e847e0005001e818e000100",
		},
		{
			"0000e0208c3b3ed3aa778eaecdcbe91dae57197ce1baa0d7c33e86d00d0100000000000079ffca6c6b36348c306234dee2fe47bafd76df7e70c95cbdff3efeb81e5abe71ea88b860fcff031a45722027",
			[]string{
				"010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff240381841e0c2074657374206d696e65722012097573200909200902825401fa4184010000ffffffff02efbc97000000000017a9140d2eb00a31486c91e3dbefa13ac714e236390dad870000000000000000266a24aa21a9ed02ee31a4ff032e606a5bc1af454ddca6695a1261a69d4ddb24d6dd10cb6d3fcd0120000000000000000000000000000000000000000000000000000000000000000000000000",
				"02000000000101761d35946ece6a53b79380119ccda626a8efd5caee724d71d81f404f5a33003f00000000171600145074935eaaf3cc1f04acc64c2c4f88737ff17896feffffff02f0f20f000000000017a9147e96bcc24d343e35f857f593eca765ccdf200b17872172b38a0000000017a914647dbca76f3d3426a564361c9539aae810752e4d870247304402204e80eb98037ec577a88a4712e6b3ea81eb23541052ecd595211edce78034fbd8022048bc0507ca70261810906c818e761d6a53cb69ad11c2c19ac554047622e88f85012103521bdcb10ea983094184c8b8dc49698541d63c45d91d90771b75434d568c80fb7f841e00",
				"02000000000101b8ff4b6851eee40c958c58916604c15533e7fb4a64a9ec509a55d10167b64b290000000000feffffff023b1ccdd700000000160014484f0aa3ea9b74beb476be62936df51c2fc59b99e84610000000000016001448717afa6934ad4da8e5e827ca061d11e41f78e602473044022079d02fd7cca6aa2f3e8860987b3fae269677ea97bc1c79553d1e53b6d5329454022040ecf6545258d0702713f46d4bafe1f1eb05c1dec983b84982d1cf2c807195dd0121034120b65994ee9788e450312b312d1fcef975c0f34dadd0fbe8ed01c9c61633a880841e00",
				"0100000001e7742bc7ac999bd4b0832809534a5965cf5b53abdc75e3f840702be92b1f82dd030000006a47304402202fe1e5defb67549a2f2a8b7d754a80866583f348e2aae9a61ece13b0842b16ca0220311eacc4c8bf859ba2ad76b7093caa366cff4b4ec66fb6e14708915cad3a04e40121037435c194e9b01b3d7f7a2802d6684a3af68d05bbf4ec8f17021980d777691f1dfdffffff040000000000000000536a4c5058365b8ca81abb972accd336d8f05be0c843b5486fb72528877a5f50e898b14d3b39e011a46ff8b0ad5a933ab2caee547240cb41e7da3cf69d5f364f4ba76e06796dc5001e847e0005001e818e00010010270000000000001976a914000000000000000000000000000000000000000088ac10270000000000001976a914000000000000000000000000000000000000000088acc1e09a04000000001976a914ba27f99e007c7f605a8305e318c1abde3cd220ac88ac00000000",
				"0100000001d6541ff19b573a4742925a56552e608fc827a149e412e30018e64fc39c6ebcbb010000008b483045022100c9a4d2502164a78caefd4d3b1deac72bdf418636e2cf16bea7a051438bec8499022070dcebad8a71a1acc6fea62512ae88a9e7901d31cbeca4a1d1af10640eba538c0141045a5ddc925295b71bafbe56bf4c10e1c1bc7c3a2bf5116b72f5dd202bccc032955afc5191f626284508072d397fd0fde700ae6feb2a35c1c391b12971960e6df6ffffffff03ee050000000000001600140fb58dc4fc27d579fd59cd18d3b44f8b5df1b47b2b2be30b000000001976a914d9ea351605b36fc3a967d790132230eb7eced36688ac0000000000000000256a2302000fa26dbf437f2811124e8395d532f969f2ee83a6d0542e2f5798ce37f267f2fdaa00000000",
				"01000000000101614b0cdbc00644c6e8bb016ab669acafacab0279c4df379440cf225ed1fd2c8a0100000000ffffffff02407e05000000000017a914dc75fc89f54f9618ee4fb5ef538c3baa46adf7458719480800000000001600146bee0d3f361510aae3e9d5014f8f91db21342506024830450221009315bcf6106e8666ea4dea43f09a81550f4b481b67e20375a3acb2678de500e402201d8a79065987b1cad8d032d4e58652d6258d4e3f53b7d15b2a78029c77f595100121035ad29a75b4561eb39c4f81a6f860d032385ef56717f3b896322c9a611cb40ed300000000",
				"02000000000101bfb1f07e1055a1ec06c045322ce1d8354d59f10fea0ed3294bd12954150f404f0100000000fdffffff02142b020000000000160014f2bb4458b593dfd6d96d1474469e240121ccfc0540420f0000000000160014206f30d19f2d39dfeda1933e07a62e38ce380da10247304402202576f7adcca78e963366c7b16c3732f3401ccc42e1726ec1e675cecd1928591a022028b60a4ab7de85112252aacaca74f3caf4927a2932c719cecb008194a51a45ad012102b65c0805712db5158e1554297fe51e468e8bbb25464d3fed183d8e7b969b539e7d841e00",
				"020000000001032c5faf7fc3751f79f93ed389db1af6394fffede942d51fcec61d15bb5829302d01000000171600147af2e87a43f5370b17bb0bd0ea18849e123ba61afeffffffad6d0fe044f3440347fa07e2fe07fac2d1a585f0a66b9a23265d968667aa21ec0000000017160014b5d5a81ba37d47ad5fe276b23f3973f56b70a459feffffff2b564e1dd88cbddbd79ba2b5467f64cbbec439c008738906e361f1b2f5eb0fe40000000000feffffff02801a060000000000160014407d31e030df0f64a73f142a0feecd4875a0626c32810f00000000001600149cffa54d75caa5cb3897efde3ef9017d593e203a0247304402200bf2fcd9dec73082763bef3999f531f9b6b5f707535596244623f874973de610022013ac74dd9a5b6eb7300e35bea44546f62b726ec4b277342920d2378df34361a501210253d9126db8349ae5550fadd02ebe612fea2dd239ba32f1029342d407d4ab6fba02473044022048b1d568ee52de84ca71dddab0503e1e8f4829a6478a1d69baeab8896ad5c9230220757e3255cf415e067f843396e5226ff6ed8facb45013773e18aa9bb0243791f50121038a51c7db5ab7377e8d5f9a06be6b6b46c900df51af9495a79f3ef01a46c71f4902473044022005030ff6d9539774eb1a03eb3ae06e73c15011404076cac08dc2fbd4d54b67a0022052d902113c9ed40c326bea976d3f0ce475fc3a575ea55e9bf158d410fc842f46012103c2f8ebfebd41bfa467ab63f3ef7e734c64ee38a46fc2f4207f3aca935bd170be7f841e00",
				"010000000001010f90038c463b5f6fc0593ec8110951c665052b2631e5c0978f65e48706d195b30100000000ffffffff02b5264408000000001600147fa2f0e860094b79fda8ac58245e2c9c5c1fbb1f8c5a0000000000001976a9141d7cb15cc4393354b4007e97161c58c1e16f0c0d88ac02473044022030aa6f4718e6dc75e0448b6de4c4d8a2a0c1c49be5c5eca1d3902b1b8a44c9da02201982eead66664a27637b46b90fc88bd80126208381700be6be327f0c8c79386d0121029459153248d701d7402b0ef3dddfde42f3867a52f172fa383746287014dd4a2e00000000",
			},
			4,
			"02000fa26dbf437f2811124e8395d532f969f2ee83a6d0542e2f5798ce37f267f2fdaa",
		},
		{
			"000000209e27bfdf755b434d57f3ce1a2045a6d1aba1832e07f209712700000000000000045e3c7c45090708ee3250519041bae1f9aa4df0cfe678f4fe63171a7c1d1edf9694b860fcff031a84f9c7b8",
			[]string{
				"020000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff1503b2841e049694b8600134ea00030000ff00000000ffffffff02f4d597000000000017a914f5fb634163aee17801523fbfaee93d4baa6cf383870000000000000000266a24aa21a9ed95d6a57229d88063d35b29112f72bab9c0c03d7f4dd9e21b20f3c7f4cc9acb060120000000000000000000000000000000000000000000000000000000000000000000000000",
				"02000000000101e82231797b226eba2bce4ebec9e0052f76f7bca5dc36a51a06adacf17952daa10000000017160014bda2810a2886a5dea011c5757c74b49d7d4354defeffffff02ba511b00000000001976a91439f848db5fa3d62bed2960e8d171faab95ed299a88ac8de79d870000000017a914af4097ba2f04bfa21becae02d3581f9d00ca7a1187024730440220341fe850161413821f86327ab6a5ccfdec4af1b426569cddb9ac53489843456c02202b042e702545d0f2e6324e8f278417ee210ece39d334ddf54f8474af9b6fd7be01210293d3a3e5160d74990a17317c7b21aef661f245f69cbd6aeabc8f1432fdb7cdc9b1841e00",
				"02000000000101a6837fadd3f7a3395c8554e777343e48e046b34bd509615638277f4f10de5a780100000000feffffff02a49a05000000000017a9140b25fcf671a0e8f8be90e574d93d7eeb10bb4df9871cca45d60000000017a914eb579dca4d3be004509d6ccadcdee8e5dfd589b787024730440220325a54e10283495151baac99693acc02b0dfafc2012b8837c4945831003e593802201705713dd0c183ef1f80676a50e9e87fd157626e6ef1320794c101fb4e2a3750012102419b7dbbdba1486ce4f16dd65aa933462996d41d44904578c3ac0ecd5eb729b8b1841e00",
				"02000000000102d2fe7f7a75c381e2e7ec14f53f12ebe4d46a7762625555887a537fe0d4fcc72d0100000017160014322c1699d34b7d89c0bd1e7c7ce9687aa6eddb8201000000d3258e02da986d4bb4eb798582c32260d8433afeabf8c0648f1435263d63a12c0100000017160014322c1699d34b7d89c0bd1e7c7ce9687aa6eddb820100000001b28301000000000017a914334358eb21e5e497a50a2b797ad1a23a1a1fb45687024730440220458ba20f33b789393143116bfd41beef575ace5810bb53e1f7f9bcb736c55fcc02200ec6c7f951b3e77e7c3d9b4a6f058378d71c0cec44a162de5508dc90d019dc4f012102f17e4ed57d3570338a6f956d1e00bdd639dbd87f1b36a45a6a25ed75c01a988e024730440220417dc98167dffe9f8bce08a061268c7aca803c959207aaf1d04fc2f25092374402201cff35b71f5e3a2d8b2e9befc4ca410e7a65c79a93ee8ee99d10c7090255e85c012102f17e4ed57d3570338a6f956d1e00bdd639dbd87f1b36a45a6a25ed75c01a988e00000000",
				"020000000001011a78ddae1c56014a78e0d15d432d48e4f34922ef3bd9d6302ee7b9b6cdc0b6960000000000feffffff030000000000000000226a2000d2833d4071a235a7bf5aa87999e78587946aa89882ce8c118c97094b6e9defc011000000000000160014dd06c079ec009f9d758056c24981e92c60b4e04b3547000000000000160014971494f1a0eddce971754b93b65cf4e855cba03002473044022072c271dfd9077760ed390418cfd2bbc1de32684e81104334371462e561d9183f022027bf6885cc72d30d02debb0ac523a04cfcdc585f03e18a319cb0f0f9f2ed854c0121026925411a455fb52a628259a5c5833c82c5a6066a93a125668f8eff4eeace2d8f00000000",
				"02000000000101555fe667261847048d1a075eff879ba526ceccb3446c375a76a48dbd0bfd38710100000000feffffff020000000000000000226a20e81a786c066da715456ddf31d6d80a8689620a0bd722b18532ca481dc58ef60381040000000000001600142b4d3909e075fc496ba6fbc051fbe4c1671475670247304402205f2ea5c0092cde961afb500bb4446f7e6c1b134efa3774b8bc22f71699062f5d02207d5f85cca2e62ad86d32d1991dee4c0f75f9cae68f5ac227f49f26c63d4c15aa012102f15e9d9d6e35e805888aedd8b6469cf2607b784e39cb2bf90d7db2fc242a666700000000",
				"02000000000101b3354159353c9b86b9b3f5cf14699d1d92a12c4ff0430c154b908cb5073ca0430100000017160014a0642eba4abe84f01ac443727a92a83aacbd3d4cfeffffff02dbe807010000000017a9143f9055d003aedff712dffecdb90d8c590132261787a08601000000000017a914c6953606f7d751d8c1a956c888ce96ed97d7e09b870247304402203604f2167529c724bfe57f893652106ff71d81979b95c531028099c100c9746002202c12c7c0efed867556cb14bb310e7cc0b01c13037cb824feab00247482c1e7e1012102dcfd099718021f3792dc9222e65882722210982ba38a905b23236cbfd5e97bacb0841e00",
			},
			4,
			"00d2833d4071a235a7bf5aa87999e78587946aa89882ce8c118c97094b6e9def",
		},
		{
			"000000209e27bfdf755b434d57f3ce1a2045a6d1aba1832e07f209712700000000000000045e3c7c45090708ee3250519041bae1f9aa4df0cfe678f4fe63171a7c1d1edf9694b860fcff031a84f9c7b8",
			[]string{
				"020000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff1503b2841e049694b8600134ea00030000ff00000000ffffffff02f4d597000000000017a914f5fb634163aee17801523fbfaee93d4baa6cf383870000000000000000266a24aa21a9ed95d6a57229d88063d35b29112f72bab9c0c03d7f4dd9e21b20f3c7f4cc9acb060120000000000000000000000000000000000000000000000000000000000000000000000000",
				"02000000000101e82231797b226eba2bce4ebec9e0052f76f7bca5dc36a51a06adacf17952daa10000000017160014bda2810a2886a5dea011c5757c74b49d7d4354defeffffff02ba511b00000000001976a91439f848db5fa3d62bed2960e8d171faab95ed299a88ac8de79d870000000017a914af4097ba2f04bfa21becae02d3581f9d00ca7a1187024730440220341fe850161413821f86327ab6a5ccfdec4af1b426569cddb9ac53489843456c02202b042e702545d0f2e6324e8f278417ee210ece39d334ddf54f8474af9b6fd7be01210293d3a3e5160d74990a17317c7b21aef661f245f69cbd6aeabc8f1432fdb7cdc9b1841e00",
				"02000000000101a6837fadd3f7a3395c8554e777343e48e046b34bd509615638277f4f10de5a780100000000feffffff02a49a05000000000017a9140b25fcf671a0e8f8be90e574d93d7eeb10bb4df9871cca45d60000000017a914eb579dca4d3be004509d6ccadcdee8e5dfd589b787024730440220325a54e10283495151baac99693acc02b0dfafc2012b8837c4945831003e593802201705713dd0c183ef1f80676a50e9e87fd157626e6ef1320794c101fb4e2a3750012102419b7dbbdba1486ce4f16dd65aa933462996d41d44904578c3ac0ecd5eb729b8b1841e00",
				"02000000000102d2fe7f7a75c381e2e7ec14f53f12ebe4d46a7762625555887a537fe0d4fcc72d0100000017160014322c1699d34b7d89c0bd1e7c7ce9687aa6eddb8201000000d3258e02da986d4bb4eb798582c32260d8433afeabf8c0648f1435263d63a12c0100000017160014322c1699d34b7d89c0bd1e7c7ce9687aa6eddb820100000001b28301000000000017a914334358eb21e5e497a50a2b797ad1a23a1a1fb45687024730440220458ba20f33b789393143116bfd41beef575ace5810bb53e1f7f9bcb736c55fcc02200ec6c7f951b3e77e7c3d9b4a6f058378d71c0cec44a162de5508dc90d019dc4f012102f17e4ed57d3570338a6f956d1e00bdd639dbd87f1b36a45a6a25ed75c01a988e024730440220417dc98167dffe9f8bce08a061268c7aca803c959207aaf1d04fc2f25092374402201cff35b71f5e3a2d8b2e9befc4ca410e7a65c79a93ee8ee99d10c7090255e85c012102f17e4ed57d3570338a6f956d1e00bdd639dbd87f1b36a45a6a25ed75c01a988e00000000",
				"020000000001011a78ddae1c56014a78e0d15d432d48e4f34922ef3bd9d6302ee7b9b6cdc0b6960000000000feffffff030000000000000000226a2000d2833d4071a235a7bf5aa87999e78587946aa89882ce8c118c97094b6e9defc011000000000000160014dd06c079ec009f9d758056c24981e92c60b4e04b3547000000000000160014971494f1a0eddce971754b93b65cf4e855cba03002473044022072c271dfd9077760ed390418cfd2bbc1de32684e81104334371462e561d9183f022027bf6885cc72d30d02debb0ac523a04cfcdc585f03e18a319cb0f0f9f2ed854c0121026925411a455fb52a628259a5c5833c82c5a6066a93a125668f8eff4eeace2d8f00000000",
				"02000000000101555fe667261847048d1a075eff879ba526ceccb3446c375a76a48dbd0bfd38710100000000feffffff020000000000000000226a20e81a786c066da715456ddf31d6d80a8689620a0bd722b18532ca481dc58ef60381040000000000001600142b4d3909e075fc496ba6fbc051fbe4c1671475670247304402205f2ea5c0092cde961afb500bb4446f7e6c1b134efa3774b8bc22f71699062f5d02207d5f85cca2e62ad86d32d1991dee4c0f75f9cae68f5ac227f49f26c63d4c15aa012102f15e9d9d6e35e805888aedd8b6469cf2607b784e39cb2bf90d7db2fc242a666700000000",
				"02000000000101b3354159353c9b86b9b3f5cf14699d1d92a12c4ff0430c154b908cb5073ca0430100000017160014a0642eba4abe84f01ac443727a92a83aacbd3d4cfeffffff02dbe807010000000017a9143f9055d003aedff712dffecdb90d8c590132261787a08601000000000017a914c6953606f7d751d8c1a956c888ce96ed97d7e09b870247304402203604f2167529c724bfe57f893652106ff71d81979b95c531028099c100c9746002202c12c7c0efed867556cb14bb310e7cc0b01c13037cb824feab00247482c1e7e1012102dcfd099718021f3792dc9222e65882722210982ba38a905b23236cbfd5e97bacb0841e00",
			},
			5,
			"e81a786c066da715456ddf31d6d80a8689620a0bd722b18532ca481dc58ef603",
		},
	}

	for i, test := range tests {

		headerBytes, _ := bbn.NewBTCHeaderBytesFromHex(test.header)

		var transactionBytes [][]byte

		for _, tx := range test.transactions {
			tb, _ := hex.DecodeString(tx)

			transactionBytes = append(transactionBytes, tb)
		}

		branch, _ := btcctypes.CreateProofForIdx(transactionBytes, uint(test.opReturnTransactionIdx))

		opReturnTx := transactionBytes[test.opReturnTransactionIdx]

		var cProof []byte

		for _, h := range branch {
			cProof = append(cProof, h.CloneBytes()...)
		}

		p, err := btcctypes.ParseProof(
			opReturnTx,
			uint32(test.opReturnTransactionIdx),
			cProof,
			&headerBytes,
			btcchaincfg.TestNet3Params.PowLimit)

		if err != nil {
			t.Errorf("Test failed due to: %v", err)
		}

		expectedOpReturn, _ := hex.DecodeString(test.expectedOpReturnData)

		if !bytes.Equal(expectedOpReturn, p.OpReturnData) {
			t.Errorf("Test %d does not contain expected op return data", i)
		}
	}
}
