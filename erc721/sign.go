package erc721

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/inwecrypto/sha3"
)

const (
	ERC721_balanceOf           = "balanceOf(address)"
	ERC721_totalSupply         = "totalSupply()"
	ERC721_transfer            = "transfer(address,uint256)"
	ERC721_decimals            = "decimals()"
	ERC721_name                = "name()"
	ERC721_symbol              = "symbol()"
	ERC721_ownerOf             = "ownerOf(uint256)"
	ERC721_approve             = "approve(address,uint256)"
	ERC721_setApprovalForAll   = "setApprovalForAll(address,bool)"
	ERC721_getApprovedAddress  = "getApprovedAddress(uint256)"
	ERC721_isApprovedForAll    = "isApprovedForAll(address,address)"
	ERC721_takeOwnership       = "takeOwnership(uint256)"
	ERC721_tokenOfOwnerByIndex = "tokenOfOwnerByIndex(address,uint256)"
	ERC721_tokenMetadata       = "tokenMetadata(uint256)"
	ERC721_tokensOf            = "tokensOf(address)"
	ERC721_exists              = "exists(uint256)"
	ERC721_setAssetHolder      = "setAssetHolder(address,uint256)"
	ERC721_transferFrom        = "transferFrom(address,address,uint256)"
	ERC721_isAuthorized        = "isAuthorized(address,uint256)"
	ERC721_description         = "description()"
	ERC721_safeTransferFrom    = "safeTransferFrom(address,address,uint256,bytes)"
)

// Method/Event id
var (
	Method_transfer            = SignABI(ERC721_transfer)
	Method_balanceOf           = SignABI(ERC721_balanceOf)
	Method_decimals            = SignABI(ERC721_decimals)
	Method_totalSupply         = SignABI(ERC721_totalSupply)
	Method_name                = SignABI(ERC721_name)
	Method_symbol              = SignABI(ERC721_symbol)
	Method_ownerOf             = SignABI(ERC721_ownerOf)
	Method_approve             = SignABI(ERC721_approve)
	Method_setApprovalForAll   = SignABI(ERC721_setApprovalForAll)
	Method_getApprovedAddress  = SignABI(ERC721_getApprovedAddress)
	Method_isApprovedForAll    = SignABI(ERC721_isApprovedForAll)
	Method_takeOwnership       = SignABI(ERC721_takeOwnership)
	Method_tokenOfOwnerByIndex = SignABI(ERC721_tokenOfOwnerByIndex)
	Method_tokenMetadata       = SignABI(ERC721_tokenMetadata)
	Method_tokensOf            = SignABI(ERC721_tokensOf)
	Method_exists              = SignABI(ERC721_exists)
	Method_setAssetHolder      = SignABI(ERC721_setAssetHolder)
	Method_transferFrom        = SignABI(ERC721_transferFrom)
	Method_isAuthorized        = SignABI(ERC721_isAuthorized)
	Method_description         = SignABI(ERC721_description)
)

// SignABI sign abi string
func SignABI(abi string) string {
	hasher := sha3.NewKeccak256()
	hasher.Write([]byte(abi))
	data := hasher.Sum(nil)

	return hex.EncodeToString(data[0:4])
}

// BalanceOf create erc20 balanceof abi string
func BalanceOf(address string) string {
	address = strings.Trim(address, "0x")

	return fmt.Sprintf("0x%s%s", Method_balanceOf, packNumeric(address, 32))
}

// GetDecimals .
func GetDecimals() string {
	return fmt.Sprintf("0x%s", Method_decimals)
}

func GetName() string {
	return fmt.Sprintf("0x%s", Method_name)
}

func GetSymbol() string {
	return fmt.Sprintf("0x%s", Method_symbol)
}

func OwnerOf(value string) string {
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)
	data := fmt.Sprintf("%s%s", Method_ownerOf, value)

	return data
}

func TokensOf(address string) string {
	address = strings.Trim(address, "0x")

	return fmt.Sprintf("0x%s%s", Method_tokensOf, packNumeric(address, 32))
}

func SetAssetHolder(to string, value string) string {
	to = packNumeric(strings.TrimPrefix(to, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("%s%s%s", Method_setAssetHolder, to, value)
}

func GetTokenMetadata(value string) string {
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)
	return fmt.Sprintf("%s%s", Method_tokenMetadata, value)
}

func Approve(to string, value string) string {
	to = packNumeric(strings.TrimPrefix(to, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("%s%s%s", Method_approve, to, value)
}

func packNumeric(value string, bytes int) string {
	value = strings.TrimSuffix(value, "0x")

	chars := bytes * 2

	n := len(value)
	if n%chars == 0 {
		return value
	}
	return strings.Repeat("0", chars-n%chars) + value
}

// Transfer .
func Transfer(to string, value string) string {
	to = packNumeric(strings.TrimPrefix(to, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("%s%s%s", Method_transfer, to, value)
}

func TransferFrom(from, to string, value string) string {
	from = packNumeric(strings.TrimPrefix(from, "0x"), 32)
	to = packNumeric(strings.TrimPrefix(to, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("%s%s%s%s", Method_transferFrom, from, to, value)
}
