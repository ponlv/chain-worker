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
			"aaa": job.MintNFT,
		})

	<-forever
}
