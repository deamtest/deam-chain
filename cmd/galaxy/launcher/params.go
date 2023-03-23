package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// mainnet
		"enode://9bf6b858c9de58cf063cd4bdcd8fb5e4ff19371f14624a3f65bdaa24764ed32f192747deba9a99a69cf91692cec823aee74c4be7353c3a1373722bccfb793954@88.119.171.113:18887",
		"enode://32ab96f5ce27a6577e3673d68d93f256636d4ece9102954a19bf439ec85f548d73748aa392024cc7ba22dcc741f96af6cd10b1324cd8a9c7873ed8019031e8ec@88.119.175.173:18887",
		"enode://3a060eddef6f309c2efb57d7f7fdd8fe72c20cf97717229cee59a89aeb5c3265b31c09a62e8ed0bebfbe3bf0e3158524f3b6e16d2c0dac39f448c377e0dce1d7@88.119.175.211:18887",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
