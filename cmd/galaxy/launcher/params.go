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
		"enode://50eaf7764633801779d8177f11570435222f8cc36e062e58e6d1762c77d9f8b9de6abee4096f63e831fc4bd9f44347a9cca1b8c6bde2b0c7b736299c8fec88b1@185.74.222.233:18887",
		"enode://@185.74.222.235:18887",
		"enode://@194.156.98.9:18887",
		"enode://@194.156.98.10:18887",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
