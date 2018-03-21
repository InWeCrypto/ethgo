package main

import (
	"fmt"

	"github.com/inwecrypto/ethgo/rpc"

	"github.com/inwecrypto/ethgo/erc20"
)

var client *rpc.Client

func init() {
	client = rpc.NewClient("http://47.52.158.99:8545")
}

func main() {
	accoutState, err := client.GetBalance("0xf4cc4154a4987f8784064468c3c6b21f0d0cdd64")

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(accoutState)

	state := &rpc.CallSite{
		From:     "0xf4cc4154a4987f8784064468c3c6b21f0d0cdd64", // from
		To:       "0x9b7929b142dddc08b889c146340c872cf8d6de71", // smart comtract address
		Value:    "0x0",
		GasPrice: "0x0",
		Gas:      "0x0",
		Data:     erc20.BalanceOf("0xf4cc4154a4987f8784064468c3c6b21f0d0cdd64"), // comtract function
	}

	fmt.Println(state)

	val, err := client.Call(state)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(val)

	intval, err := rpc.ReadBigint(val)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(intval)
}
