package main

import (
	"fmt"
	"reflect"

	"math/big"
	"runtime"

	"github.com/carlescere/scheduler"
	"github.com/ethereum/go-ethereum/common/hexutil"

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
		msgs, err := shh.GetFilterMsgs(filterID)
		if err != nil {
			panic(err)
		}

		if res, ok := msgs.Result.([]interface{}); ok {
			for _, msg := range res {
				if msgMap, ok := msg.(map[string]interface{}); ok {
					payload, err := hexutil.Decode(msgMap["payload"].(string))
					if err != nil {
						panic(err)
					}
					fmt.Println(string(payload))
					for k, v := range msgMap {
						fmt.Printf("key : %s val : %s\n", k, v)
					}
				}
				//r := msgs.Result.(map[string]interface{})["payload"]
			}

		} else {
			fmt.Println("not")
			fmt.Println(reflect.TypeOf(msgs.Result))
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
	_, err = shh.AsymPost(sender, pubKey, "0xdeadbeef", "Heloo mesg", big.NewInt(7))
	if err != nil {
		panic(err)
	}
	//fmt.Println(res)
	// Keep the program from not exiting.
	runtime.Goexit()
}
