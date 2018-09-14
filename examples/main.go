package main

import (
	"fmt"
	"math/big"
	"runtime"

	"github.com/enkhalifapro/go-web3/utils"

	"github.com/carlescere/scheduler"

	"github.com/enkhalifapro/go-web3/dto"
	"github.com/enkhalifapro/go-web3/providers"
	"github.com/enkhalifapro/go-web3/shh"
)

func main() {
	provider := providers.NewHTTPProvider("127.0.0.1:8545", 10, false)
	shh := shh.NewSHH(provider)

	// 1- create a new keypair
	keyID, err := shh.NewKeyPair()
	if err != nil {
		panic(err)
	}

	//keyID = "a289d0fbb21351ce4a780d081c5eda6e8f1f52dfb564b6fe9ee977643e60191a"

	// 2- create a message filter
	filterID, err := shh.NewMsgFilter(&dto.SHHSubscribeParam{
		PrivateKeyID: keyID,
		Topics:       []string{"0xdeadbeef"},
	})
	if err != nil {
		panic(err)
	}

	getMsgs := func() {
		msgs := shh.GetFilterMsgs(filterID)
		if len(msgs) > 0 {
			for _, msg := range msgs {
				fmt.Println(utils.DecodeHex(msg.Payload))
			}
		}
	}

	scheduler.Every(1).Seconds().Run(getMsgs)

	// 3- get the public key (which is needed by the sender)
	pubKey, err := shh.GetPublicKey(keyID)
	if err != nil {
		panic(err)
	}

	// 4- send a message to subscriber
	sender, err := shh.NewKeyPair()
	for index := 0; index < 1000; index++ {
		_, err = shh.AsymPost(sender, pubKey, "0xdeadbeef", fmt.Sprintf("Hello %v", index), big.NewInt(7))
		if err != nil {
			panic(err)
		}
	}

	// Keep the program from not exiting.
	runtime.Goexit()
}
