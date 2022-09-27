package main

import (
	"chain-worker/configs"
	"chain-worker/modules/job"
	"github.com/ponlv/go-kit/ethereum"
	"github.com/ponlv/go-kit/ethereum/ethworker"
)

func main() {
	configs.InitConfig()

	ethereum.NewClient(configs.SchemaConfig.Chain.Url)

	forever := make(chan bool)

	ethworker.ListenEvent(
		configs.SchemaConfig.Chain.WSS,
		[]string{""}, // TODO: add your smart contract address here
		1123,         // TODO: add your smart contract block here
		map[string]ethworker.ChainEventHandler{
			"0xc5c2f7a3cc1db796eab9ab3ac964ecf4689f447ddbfbdafbb3b9547c9d1883ff": job.MintNFT,
		})

	<-forever
}
