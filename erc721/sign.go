package erc721

import (
	"encoding/hex"
	"fmt"
	. "github.com/ximenyan/ethgo/contract"
	"math/big"
)

const (
	ownerOf             = "ownerOf(uint256)"
	setApprovalForAll   = "setApprovalForAll(address,bool)"
	getApprovedAddress  = "getApprovedAddress(uint256)"
	isApprovedForAll    = "isApprovedForAll(address,address)"
	takeOwnership       = "takeOwnership(uint256)"
	tokenOfOwnerByIndex = "tokenOfOwnerByIndex(address,uint256)"
	tokenMetadata       = "tokenMetadata(uint256)"
	tokensOf            = "tokensOf(address)"
	exists              = "exists(uint256)"
	setAssetHolder      = "setAssetHolder(address,uint256)"
	isAuthorized        = "isAuthorized(address,uint256)"
	description         = "description()"

	// DecentraLand
	DecentraLand_decodeTokenId = "decodeTokenId(uint256)"
	DecentraLand_encodeTokenId = "encodeTokenId(int256,int256)"
	DecentraLand_landData      = "landData(int256,int256)"
	DecentraLand_landOf        = "landOf(address)"
	DecentraLand_transferLand  = "transferLand(int256,int256,address)"
	DecentraLand_ownerOfLand   = "ownerOfLand(int256,int256)"

	// RedPacket
	RedPacket_newRedPacket           = "newRedPacket(uint256,address,address,uint256,uint256,uint256)"
	RedPacket_openMany               = "openMany(uint256,address[],uint256,bool)"
	RedPacket_setTaxCost             = "setTaxCost(uint256,uint256)"
	RedPacket_changeWallet           = "changeWallet(address)"
	RedPacket_changeMaxCount         = "changeMaxCount(uint256)"
	RedPacket_getRedPacketOpenDetail = "getRedPacketOpenDetail(uint256)"
	RedPacket_getRedPacketStatus     = "getRedPacketStatus(uint256)"
	RedPacket_sendEther              = "sendEther(uint256)"
	RedPacket_getTaxCost             = "getTaxCost()"
	RedPacket_maxCount               = "maxCount()"
	RedPacket_addAdmin               = "addAdmin(address)"
	RedPacket_delAdmin               = "delAdmin(address)"
	RedPacket_changeGatherValue      = "changeGatherValue(uint256)"
)

// Method/Event id
var (
	Method_ownerOf             = SignABI(ownerOf)
	Method_setApprovalForAll   = SignABI(setApprovalForAll)
	Method_getApprovedAddress  = SignABI(getApprovedAddress)
	Method_isApprovedForAll    = SignABI(isApprovedForAll)
	Method_takeOwnership       = SignABI(takeOwnership)
	Method_tokenOfOwnerByIndex = SignABI(tokenOfOwnerByIndex)
	Method_tokenMetadata       = SignABI(tokenMetadata)
	Method_tokensOf            = SignABI(tokensOf)
	Method_exists              = SignABI(exists)
	Method_setAssetHolder      = SignABI(setAssetHolder)
	Method_isAuthorized        = SignABI(isAuthorized)
	Method_description         = SignABI(description)
	Method_transferLand        = SignABI(DecentraLand_transferLand)
)


func GetDescription() string {
	return fmt.Sprintf("0x%s", Method_description)
}

func OwnerOf(value string) string {
	value = PackNumeric(value, 32)
	return fmt.Sprintf("0x%s%s", Method_ownerOf, value)
}

func TokensOf(address string) string {
	address = PackNumeric(address, 32)

	return fmt.Sprintf("0x%s%s", Method_tokensOf, address)
}

func SetAssetHolder(to string, value string) ([]byte, error) {
	to = PackNumeric(to, 32)
	value = PackNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s", Method_setAssetHolder, to, value)

	return hex.DecodeString(data)
}

func GetTokenMetadata(value string) string {
	value = PackNumeric(value, 32)
	return fmt.Sprintf("0x%s%s", Method_tokenMetadata, value)
}

func TransferLand(to string, x, y string) ([]byte, error) {
	to = PackNumeric(to, 32)
	x = PackNumeric(x, 32)
	y = PackNumeric(y, 32)

	data := fmt.Sprintf("%s%s%s%s", SignABI(DecentraLand_transferLand), x, y, to)

	return hex.DecodeString(data)
}

func IsExists(value string) string {
	value = PackNumeric(value, 32)

	return fmt.Sprintf("0x%s%s", Method_exists, value)
}

func TokenOfOwnerByIndex(adress string, value string) string {
	adress = PackNumeric(adress, 32)
	value = PackNumeric(value, 32)

	return fmt.Sprintf("0x%s%s%s", Method_tokenOfOwnerByIndex, adress, value)
}

func TakeOwnership(value string) ([]byte, error) {
	data := fmt.Sprintf("%s%s", Method_takeOwnership, value)

	return hex.DecodeString(data)
}

func DecodeTokenId(value string) string {
	value = PackNumeric(value, 32)

	return fmt.Sprintf("0x%s%s", SignABI(DecentraLand_decodeTokenId), value)
}

func EncodeTokenId(x, y string) string {
	x = PackNumeric(x, 32)
	y = PackNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", SignABI(DecentraLand_encodeTokenId), x, y)
}

func LandData(x, y string) string {
	x = PackNumeric(x, 32)
	y = PackNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", SignABI(DecentraLand_landData), x, y)
}

func Description() string {
	return fmt.Sprintf("0x%s", Method_description)
}

func LandOf(address string) string {
	address = PackNumeric(address, 32)

	return fmt.Sprintf("0x%s%s", SignABI(DecentraLand_landOf), address)
}

func OwnerOfLand(x, y string) string {
	x = PackNumeric(x, 32)
	y = PackNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", SignABI(DecentraLand_ownerOfLand), x, y)
}

func TaxCost() string {
	return "0x" + SignABI(RedPacket_getTaxCost)
}

func MaxCount() string {
	return "0x" + SignABI(RedPacket_maxCount)
}

func SetTaxCost(min, max string) ([]byte, error) {
	data := SignABI(RedPacket_setTaxCost) + PackNumeric(min, 32) + PackNumeric(max, 32)

	return hex.DecodeString(data)
}

func ChangeWallet(address string) ([]byte, error) {
	data := SignABI(RedPacket_changeWallet) + PackNumeric(address, 32)

	return hex.DecodeString(data)
}

func ChangeMaxCount(value string) ([]byte, error) {
	data := SignABI(RedPacket_changeMaxCount) + PackNumeric(value, 32)

	return hex.DecodeString(data)
}

func GetRedPacketStatus(value string) string {
	data := "0x" + SignABI(RedPacket_getRedPacketStatus) + PackNumeric(value, 32)

	return data
}

func GetRedPacketOpenDetail(value string) string {
	data := "0x" + SignABI(RedPacket_getRedPacketOpenDetail) + PackNumeric(value, 32)

	return data
}

func NewRedPacket(tokenId, address, from string, value, count, cmd string) ([]byte, error) {

	data := SignABI(RedPacket_newRedPacket) +
		PackNumeric(tokenId, 32) +
		PackNumeric(address, 32) +
		PackNumeric(from, 32) +
		PackNumeric(value, 32) +
		PackNumeric(count, 32) +
		PackNumeric(cmd, 32)

	return hex.DecodeString(data)
}

func OpenMany(tokeId string, addresses []string, cmd string, end bool) ([]byte, error) {
	endStr := "0x0"

	if end {
		endStr = "0x1"
	}

	start := hex.EncodeToString(big.NewInt(128).Bytes())

	data := SignABI(RedPacket_openMany) +
		PackNumeric(tokeId, 32) +
		PackNumeric(start, 32) +
		PackNumeric(cmd, 32) +
		PackNumeric(endStr, 32) +
		encodeStrings(addresses)

	return hex.DecodeString(data)
}

func SendEther(value string) ([]byte, error) {
	value = PackNumeric(value, 32)

	data := SignABI(RedPacket_sendEther) + value

	return hex.DecodeString(data)
}

func ChangeRedPacketGatherValue(value string) ([]byte, error) {
	value = PackNumeric(value, 32)

	data := SignABI(RedPacket_changeGatherValue) + value

	return hex.DecodeString(data)
}

func AddRedPacketAdmin(addr string) ([]byte, error) {
	addr = PackNumeric(addr, 32)

	data := SignABI(RedPacket_addAdmin) + addr

	return hex.DecodeString(data)
}

func DelRedPacketAdmin(addr string) ([]byte, error) {
	addr = PackNumeric(addr, 32)

	data := SignABI(RedPacket_delAdmin) + addr

	return hex.DecodeString(data)
}

func encodeStrings(params []string) string {
	length := big.NewInt(int64(len(params)))

	lenStr := hex.EncodeToString(length.Bytes())

	codes := PackNumeric(lenStr, 32)

	for _, v := range params {
		codes += PackNumeric(v, 32)
	}

	return codes
}
