package uniswap

import (
	"encoding/hex"
	"github.com/inwecrypto/sha3"
)

const (
	signBalanceOf         = "balanceOf(address)"
)

// Method/Event id
var (
	BalanceOfID         = SignABI(signBalanceOf)
)

// SignABI sign abi string
func SignABI(abi string) string {
	hasher := sha3.NewKeccak256()
	hasher.Write([]byte(abi))
	data := hasher.Sum(nil)
	return hex.EncodeToString(data[0:4])
}

