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

	DecentraLand_description   = "description()"
	DecentraLand_decodeTokenId = "decodeTokenId(uint256)"
	DecentraLand_encodeTokenId = "encodeTokenId(int256,int256)"
	DecentraLand_landData      = "landData(int256,int256)"
	DecentraLand_landOf        = "landOf(address)"
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

	Method_DecentraLand_decodeTokenId = SignABI(DecentraLand_decodeTokenId)
	Method_DecentraLand_encodeTokenId = SignABI(DecentraLand_encodeTokenId)
	Method_DecentraLand_landData      = SignABI(DecentraLand_landData)
	Method_DecentraLand_description   = SignABI(DecentraLand_description)
	Method_DecentraLand_landOf        = SignABI(DecentraLand_landOf)
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
	address = packNumeric(strings.Trim(address, "0x"), 32)

	return fmt.Sprintf("0x%s%s", Method_balanceOf, address)
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

func GetDescription() string {
	return fmt.Sprintf("0x%s", Method_description)
}

func OwnerOf(value string) string {
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)
	data := fmt.Sprintf("0x%s%s", Method_ownerOf, value)

	return data
}

func TokensOf(address string) string {
	address = packNumeric(strings.Trim(address, "0x"), 32)

	return fmt.Sprintf("0x%s%s", Method_tokensOf, address)
}

func SetAssetHolder(to string, value string) string {
	to = packNumeric(strings.TrimPrefix(to, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("%s%s%s", Method_setAssetHolder, to, value)
}

func GetTokenMetadata(value string) string {
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)
	return fmt.Sprintf("0x%s%s", Method_tokenMetadata, value)
}

func Approve(to string, value string) string {
	to = packNumeric(strings.TrimPrefix(to, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("0x%s%s%s", Method_approve, to, value)
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

func IsExists(value string) string {
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("0x%s%s", Method_exists, value)
}

func TokenOfOwnerByIndex(adress string, value string) string {
	adress = packNumeric(strings.TrimPrefix(adress, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("0x%s%s%s", Method_tokenOfOwnerByIndex, adress, value)
}

func TakeOwnership(value string) string {
	return fmt.Sprintf("%s%s", Method_takeOwnership, value)
}

func DecodeTokenId(value string) string {
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	return fmt.Sprintf("0x%s%s", Method_DecentraLand_decodeTokenId, value)
}

func EncodeTokenId(x, y string) string {
	x = packNumeric(strings.TrimPrefix(x, "0x"), 32)
	y = packNumeric(strings.TrimPrefix(y, "0x"), 32)

	return fmt.Sprintf("0x%s%s%s", Method_DecentraLand_encodeTokenId, x, y)
}

func LandData(x, y string) string {
	x = packNumeric(strings.TrimPrefix(x, "0x"), 32)
	y = packNumeric(strings.TrimPrefix(y, "0x"), 32)

	return fmt.Sprintf("0x%s%s%s", Method_DecentraLand_landData, x, y)
}

func Description() string {
	return fmt.Sprintf("0x%s", Method_DecentraLand_description)
}

func LandOf(address string) string {
	address = packNumeric(strings.TrimPrefix(address, "0x"), 32)

	return fmt.Sprintf("0x%s%s", Method_DecentraLand_landOf, address)
}
