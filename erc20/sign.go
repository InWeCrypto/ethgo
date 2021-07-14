package erc20

import (
	"encoding/hex"
	"fmt"
	. "github.com/ximenyan/ethgo/contract"
)

const (
	signBalanceOf         = "balanceOf(address)"
	signTotalSupply       = "totalSupply()"
	signTransfer          = "transfer(address,uint256)"
	signTransferFrom      = "transferFrom(address,address,uint256)"
	signApprove           = "approve(address,uint256)"
	signName              = "name()"
	signSymbol            = "symbol()"
	signAllowance         = "allowance(address,address)"
	eventTransfer         = "Transfer(address,address,uint256)"
	decimals              = "decimals()"
	signTransferOwnership = "transferOwnership(address)"
)

// Method/Event id
var (
	TransferID          = SignABI(signTransfer)
	BalanceOfID         = SignABI(signBalanceOf)
	Decimals            = SignABI(decimals)
	TransferFromID      = SignABI(signTransferFrom)
	ApproveID           = SignABI(signApprove)
	TotalSupplyID       = SignABI(signTotalSupply)
	AllowanceID         = SignABI(signAllowance)
	TransferOwnershipID = SignABI(signTransferOwnership)
)


// BalanceOf create erc20 balanceof abi string
func BalanceOf(address string) string {
	address = PackNumeric(address, 32)

	return fmt.Sprintf("0x%s%s", BalanceOfID, address)
}

// GetDecimals .
func GetDecimals() string {
	return fmt.Sprintf("0x%s", Decimals)
}

func GetTotalSupply() string {
	return fmt.Sprintf("0x%s", TotalSupplyID)
}

func GetName() string {
	return "0x" + SignABI(signName)
}

func GetSignSymbol() string {
	return "0x" + SignABI(signSymbol)
}

// Transfer .
func Transfer(to string, value string) ([]byte, error) {
	to = PackNumeric(to, 32)
	value = PackNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s", SignABI(signTransfer), to, value)

	return hex.DecodeString(data)
}

// TransferFrom .
func TransferFrom(from, to string, value string) ([]byte, error) {
	from = PackNumeric(from, 32)
	to = PackNumeric(to, 32)
	value = PackNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s%s", TransferFromID, from, to, value)

	return hex.DecodeString(data)
}

// Approve .
func Approve(to string, value string) ([]byte, error) {
	to = PackNumeric(to, 32)
	value = PackNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s", ApproveID, to, value)

	return hex.DecodeString(data)
}

func Allowance(from, to string) ([]byte, error) {
	from = PackNumeric(from, 32)
	to = PackNumeric(to, 32)

	data := fmt.Sprintf("%s%s%s", AllowanceID, to, to)

	return hex.DecodeString(data)
}

func TransferOwnership(to string) ([]byte, error) {
	to = PackNumeric(to, 32)
	data := fmt.Sprintf("%s%s", TransferOwnershipID, to)

	return hex.DecodeString(data)
}
