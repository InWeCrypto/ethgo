package erc721

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/inwecrypto/ethgo"
	"github.com/inwecrypto/ethgo/keystore"
	"github.com/inwecrypto/ethgo/rpc"
	"github.com/inwecrypto/ethgo/tx"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc"
)

var key *keystore.Key
var client *rpc.Client
var jsonrpcclient *jsonrpc.RPCClient

var contractAddress string

func init() {

	// smart contract address
	contractAddress = "0x7a3ca2d04edd85602716ad7174379ef48ca809bc"

	// ganache-cli  PrivateKey
	pk := "8438a3e00668820870585d2369b0b2a58242379dad0a80cd84a1955df2db69f4"

	bs, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}

	key, err = keystore.KeyFromPrivateKey(bs)

	if err != nil {
		panic(err)
	}

	url := "http://127.0.0.1:8545"

	client = rpc.NewClient(url)

	jsonrpcclient = jsonrpc.NewRPCClient(url)
	println("private key:", key.Address)
}

// decimals .
func TestDecimals(t *testing.T) {
	deciamls, err := client.GetTokenDecimals(contractAddress)

	require.NoError(t, err)

	println("deciamls :", deciamls.Int64())
}

// balanceOf .
func TestBalanceOf(t *testing.T) {
	balance, err := client.GetTokenBalance(contractAddress, key.Address)

	require.NoError(t, err)

	println("balance :", balance.Int64())
}

// ownerOf .
func TestOwnerOf(t *testing.T) {
	value := hex.EncodeToString(big.NewInt(1).Bytes())

	data := OwnerOf(value)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	println("owner :", valstr)
}

// ReadBigint .
func ReadBigint(source string) (*big.Int, error) {
	value := big.NewInt(0)

	if source == "0x0" {
		return value, nil
	}

	source = strings.TrimPrefix(source, "0x")

	if len(source)%2 != 0 {
		source = "0" + source
	}

	data, err := hex.DecodeString(source)

	if err != nil {
		return nil, err
	}

	return value.SetBytes(data), nil
}

// tokensOf .
func TestTokensOf(t *testing.T) {
	data := TokensOf(key.Address)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	println("tokens string:", valstr)

	valstr = strings.TrimPrefix(valstr, "0x")

	for i := 0; i < len(valstr); i += 64 {
		var tokenId big.Int
		value := "0x" + valstr[i:i+64]
		err := tokenId.UnmarshalJSON([]byte(value))
		if err != nil {
			require.NoError(t, err)
		}
		println("tokenId:", tokenId.String())
	}

}

// SetAssetHolder .
func TestSetAssetHolder(t *testing.T) {
	// rand  a num
	value := hex.EncodeToString(big.NewInt(111).Bytes())

	data := SetAssetHolder(key.Address, value)

	gasLimits := big.NewInt(300000)
	gasPrice := ethgo.NewValue(big.NewFloat(20), ethgo.Shannon)
	nonce, err := client.Nonce(key.Address)

	require.NoError(t, err)

	decodeData, err := hex.DecodeString(data)

	require.NoError(t, err)

	tx := tx.NewTx(nonce, contractAddress, nil, gasPrice, gasLimits, decodeData)

	require.NoError(t, tx.Sign(key.PrivateKey))

	rawtx, err := tx.Encode()

	require.NoError(t, err)

	id, err := client.SendRawTransaction(rawtx)

	require.NoError(t, err)

	println("txid:", id)
}

func TestGetTokenMetadata(t *testing.T) {
	value := hex.EncodeToString(big.NewInt(3).Bytes())

	data := GetTokenMetadata(value)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	println("metadata :", valstr)
}

func TestTransferFrom(t *testing.T) {
	value := hex.EncodeToString(big.NewInt(3).Bytes())

	data := TransferFrom(key.Address, "0x45cdf85f75d5682df550640f6c297fc31abbbd26", value)

	gasLimits := big.NewInt(300000)
	gasPrice := ethgo.NewValue(big.NewFloat(20), ethgo.Shannon)
	nonce, err := client.Nonce(key.Address)

	require.NoError(t, err)

	decodeData, err := hex.DecodeString(data)

	require.NoError(t, err)

	tx := tx.NewTx(nonce, contractAddress, nil, gasPrice, gasLimits, decodeData)

	require.NoError(t, tx.Sign(key.PrivateKey))

	rawtx, err := tx.Encode()

	require.NoError(t, err)

	id, err := client.SendRawTransaction(rawtx)

	require.NoError(t, err)

	println("txid:", id)
}
