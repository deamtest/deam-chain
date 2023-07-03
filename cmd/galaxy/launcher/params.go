package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// mainnet

		// testnet
		"enode://121142b57506687641511c7e159f8ca355023c8b41e8cf97ccd853c45498d8d629a4522c3abee0b24b9be4c6f45dcddc249866d0dcbd51111d5e14be5293971f@65.21.251.193:18887",
		"enode://e7fd9beb9e0312370bbbaff04f7b7e07856de71b464286013d9b223f9e719a1de13fe95481fc21d1c060fd31f9a729213ffa070993e85e848122e455b8d484fb@65.109.132.178:18887",
		"enode://69339b51af2b5186a4d64c32fc3f84e5aea1bf284c484801827aea3be54693df797bc2630fd4bfea611b8d658ca5e2fb79654a9006b0283465474c37a9fc9ce9@135.181.148.22:18887",
		"enode://6308e1cf331045ba89ab88d3a91cc7157f7474c047236899dbd20f885a1d7ea7557efa77eb4c3d160ebeeaa9457e4470b40882bffee9d0cbab43656f2d2e036d@65.108.218.29:18887",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
