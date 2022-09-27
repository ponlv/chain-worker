package job

import (
	"chain-worker/modules/abi/nft"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func MintNFT(vLog types.Log) {

	contractABI, err := nft.NftMetaData.GetAbi()
	if err != nil {
		log.Println("abi.JSON", err)
		return
	} else {
		log.Println("contract abi ok")
	}

	nftEvent := nft.NftMintNFTByUser{}

	err = contractABI.UnpackIntoInterface(&nftEvent, "MintNFTByUser", vLog.Data)
	if err != nil {
		log.Println("contractAbi.Unpack", err)
		return
	}

	fmt.Println(nftEvent)

}
