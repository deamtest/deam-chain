package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// mainnet
		"enode://fbaa62a01a56b60253c80eb739ca46341617ebb295c47576f848169e005777b756cbeb1185d6eac64f2e6ce2d53e0a6269489ff260000b895ce38d9024007c81@88.119.171.113:18887",
		"enode://a2bbf5438fbf8811b700d3a6bc78f2f333abcc579e2b050cf1c30f3adf27f6663e24edd922d1b4e2e5842fe89438ca3dc7fcfd848b714e402b97620d55cbc5de@88.119.175.173:18887",
		"enode://0442c622daf1e82bfeac1b6e67daae43492c8c1ab6c8eba090f1c5a7806099288b4bbee581903a6b4e1e13df0416320fda4ae6e316d21b08b9facf952478db62@88.119.175.211:18887",

		// testnet
		"enode://098b1b46839a06cc5ffc8120db88a1143af7893c15309b967d4b9435dc1da8781c95908acb5117bc394acb8b471d9340d299e2f42205d9fc834e46c272abc7f5@185.74.222.233:18887",
		"enode://927d7aa80e2b44160173473e0cfb106715150e88903296f89ee6508a2daec0ce152c468133e0b3f4e7fd380715e56f1ded7aebda5615de4517beacb95407586f@185.74.222.235:18887",
		"enode://df35638918f6645f03e5d5682fdadf4d94c0a0c40ba883cef6705a591abe17dee33e055704ddd484f63609313a39c767ec300b39e1b409b62118ededb79a6f9e@194.156.98.9:18887",
		"enode://231df7f9e9f26b57ec6e80671c5ecc784c0c1668efede0f9fdc5f965c272a10f95acd26043768a9428fefa6a529d7bb9560606f2b510f60a0c10b16516401a1c@194.156.98.10:18887",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
