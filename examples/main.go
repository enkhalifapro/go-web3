package main

import (
	"fmt"

	"github.com/enkhalifapro/go-web3/providers"
	"github.com/enkhalifapro/go-web3/shh"
)

func main() {
	provider := providers.NewHTTPProvider("127.0.0.1:8545", 10, false)
	shh := shh.NewSHH(provider)

	v, err := shh.NewKeyPair()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
