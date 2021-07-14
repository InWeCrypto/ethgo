package uniswap

import (
	"fmt"
	c "github.com/ximenyan/ethgo/contract"
)

const (
	signGetPair       = "getPair(address,address)"
)

// Method/Event id
var (
	GetPairID         = SignABI(signGetPair)
)

func GetPair(token1,token2 string) string  {
	address1 := c.PackNumeric(token1, 32)
	address2 := c.PackNumeric(token2, 32)
	return fmt.Sprintf("0x%s%s%s", GetPairID, address1, address2)
}