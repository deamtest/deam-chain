package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// mainnet

		// testnet
		"enode://098b1b46839a06cc5ffc8120db88a1143af7893c15309b967d4b9435dc1da8781c95908acb5117bc394acb8b471d9340d299e2f42205d9fc834e46c272abc7f5@65.21.251.193:18887",
		"enode://927d7aa80e2b44160173473e0cfb106715150e88903296f89ee6508a2daec0ce152c468133e0b3f4e7fd380715e56f1ded7aebda5615de4517beacb95407586f@65.109.132.178:18887",
		"enode://df35638918f6645f03e5d5682fdadf4d94c0a0c40ba883cef6705a591abe17dee33e055704ddd484f63609313a39c767ec300b39e1b409b62118ededb79a6f9e@135.181.148.22:18887",
		"enode://231df7f9e9f26b57ec6e80671c5ecc784c0c1668efede0f9fdc5f965c272a10f95acd26043768a9428fefa6a529d7bb9560606f2b510f60a0c10b16516401a1c@65.108.218.29:18887",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
