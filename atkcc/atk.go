package atkcc

import (
	"encoding/hex"
	"fmt"
	c "github.com/ximenyan/ethgo/contract"
)

const BSC_ATK_ADDRESS = "0x5cf28E8fe45e17a15460148667e0cf4B97dE6f71"
const (
	signGetCanOutETH       = "getCanOutETH(address)"
	signATKETH = "ATKETH(address)"
)

// Method/Event id
var (
	GetCanOutETHID  = c.SignABI(signGetCanOutETH)
	ATKETHID  = c.SignABI(signATKETH)
)

func GetATKETH(_pair string) []byte {
	_pair = c.PackNumeric(_pair, 32)
	codes :=  fmt.Sprintf("%s%s", ATKETHID, _pair)
	data,_ := hex.DecodeString(codes)
	return data
}

func GetCanOutETH(_pair string) string  {
	_pair = c.PackNumeric(_pair, 32)
	return fmt.Sprintf("0x%s%s", GetCanOutETHID, _pair)
}