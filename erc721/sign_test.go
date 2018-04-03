package erc721

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/inwecrypto/ethgo/math"

	"github.com/inwecrypto/ethgo"
	"github.com/inwecrypto/ethgo/keystore"
	"github.com/inwecrypto/ethgo/rpc"
	"github.com/inwecrypto/ethgo/tx"
	"github.com/stretchr/testify/require"
)

var key *keystore.Key
var client *rpc.Client

var contractAddress string

func init() {

	// smart contract address
	//	contractAddress = "0x7a3ca2d04edd85602716ad7174379ef48ca809bc"
	contractAddress = "0xf87e31492faf9a91b02ee0deaad50d51d56d5d4d" // mainnet decentraland adress

	// ganache-cli  PrivateKey
	pk := "8438a3e00668820870585d2369b0b2a58242379dad0a80cd84a1955df2db69f4"

	bs, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}

	key, err = keystore.KeyFromPrivateKey(bs)

	//	url := "http://127.0.0.1:8545"
	url := "http://47.100.110.75:8545"

	client = rpc.NewClient(url)

	println("private key:", key.Address)
}

func TestDescription(t *testing.T) {
	data := Description()
	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	if len(valstr) < 130 {
		require.NoError(t, err)
	}

	Description, err := hex.DecodeString(valstr[130:])

	if err != nil {
		require.NoError(t, err)
	}

	println("Description :", string(Description))
}

// decimals .
func TestDecimals(t *testing.T) {
	deciamls, err := client.GetTokenDecimals(contractAddress)

	require.NoError(t, err)

	println("deciamls :", deciamls.Int64())
}

// balanceOf .
func TestBalanceOf(t *testing.T) {
	balance, err := client.GetTokenBalance(contractAddress, "0xcecbe670c11d4d28678955f23e0d2d708d79c893")

	require.NoError(t, err)

	println("balance :", balance.Int64())
}

func TestIsExists(t *testing.T) {
	var err error
	landId, b := math.ParseBig256("4423670769972200025023869896612986748907")
	if !b {
		err = errors.New("ParseBig256 failed")
	}
	require.NoError(t, err)

	value := hex.EncodeToString(landId.Bytes())

	data := IsExists(value)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	println("IsExists :", valstr)
}

func TestDecodeTokenId(t *testing.T) {
	var err error
	landId, b := math.ParseBig256("4423670769972200025023869896612986748907")
	if !b {
		err = errors.New("ParseBig256 failed")
	}
	require.NoError(t, err)

	value := hex.EncodeToString(landId.Bytes())

	data := DecodeTokenId(value)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	valstr = strings.TrimPrefix(valstr, "0x")

	println("valstr:", valstr)

	count := 1
	for i := 0; i < len(valstr); i += 64 {
		var tokenId big.Int
		value := "0x" + valstr[i:i+64]
		err := tokenId.UnmarshalJSON([]byte(value))
		if err != nil {
			require.NoError(t, err)
		}

		println("tokenId ", count, " --- ", tokenId.String())

		count++
	}
}

func TestEncodeTokenId(t *testing.T) {
	var err error
	landx, b := math.ParseBig256("12")
	if !b {
		err = errors.New("ParseBig256 failed")
	}
	require.NoError(t, err)

	x := hex.EncodeToString(landx.Bytes())

	landy, b := math.ParseBig256("115792089237316195423570985008687907853269984665640564039457584007913129639915")
	if !b {
		err = errors.New("ParseBig256 failed")
	}
	require.NoError(t, err)

	y := hex.EncodeToString(landy.Bytes())

	data := EncodeTokenId(x, y)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	println("valstr:", valstr)
	var landId big.Int
	landId.UnmarshalJSON([]byte(valstr))
	if err != nil {
		require.NoError(t, err)
	}

	println("landId ", landId.String())
}

// ownerOf .
func TestOwnerOf(t *testing.T) {
	var err error
	landId, b := math.ParseBig256("4423670769972200025023869896612986748907")
	if !b {
		err = errors.New("ParseBig256 failed")
	}

	require.NoError(t, err)

	value := hex.EncodeToString(landId.Bytes())

	data := OwnerOf(value)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	println("owner :", valstr)
}

// tokensOf .
func TestTokensOf(t *testing.T) {
	data := TokensOf("0xcecbe670c11d4d28678955f23e0d2d708d79c893")

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	valstr = strings.TrimPrefix(valstr, "0x")

	count := 1
	for i := 0; i < len(valstr); i += 64 {
		var landId big.Int
		value := "0x" + valstr[i:i+64]
		err := landId.UnmarshalJSON([]byte(value))
		if err != nil {
			require.NoError(t, err)
		}

		if count > 2 {
			println(count, " landId ", " --- ", landId.String())
		}

		count++
	}

}

func TestLandOf(t *testing.T) {
	data := LandOf("0xcecbe670c11d4d28678955f23e0d2d708d79c893")

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	valstr = strings.TrimPrefix(valstr, "0x")

	// return two arrays []int256 []int256 need to handle
	count := 1
	for i := 0; i < len(valstr); i += 64 {
		var coord big.Int
		value := "0x" + valstr[i:i+64]
		err := coord.UnmarshalJSON([]byte(value))
		if err != nil {
			require.NoError(t, err)
		}

		println(count, " coord ", " ---- ", coord.String())

		count++
	}
}

// SetAssetHolder .
func TestSetAssetHolder(t *testing.T) {
	return // ignore
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

func TestLandData(t *testing.T) {
	var err error
	landx, b := math.ParseBig256("12")
	if !b {
		err = errors.New("ParseBig256 failed")
	}
	require.NoError(t, err)

	x := hex.EncodeToString(landx.Bytes())

	landy, b := math.ParseBig256("115792089237316195423570985008687907853269984665640564039457584007913129639915")
	if !b {
		err = errors.New("ParseBig256 failed")
	}
	require.NoError(t, err)

	y := hex.EncodeToString(landy.Bytes())

	data := LandData(x, y)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	if len(valstr) < 130 {
		require.NoError(t, err)
	}

	LandData, err := hex.DecodeString(valstr[130:])

	if err != nil {
		require.NoError(t, err)
	}

	println("LandData :", string(LandData))
}

func TestGetTokenMetadata(t *testing.T) {
	var err error
	landId, b := math.ParseBig256("4423670769972200025023869896612986748907")
	if !b {
		err = errors.New("ParseBig256 failed")
	}

	value := hex.EncodeToString(landId.Bytes())

	data := GetTokenMetadata(value)

	valstr, err := client.Call(&rpc.CallSite{
		To:   contractAddress,
		Data: data,
	})

	require.NoError(t, err)

	if len(valstr) < 130 {
		require.NoError(t, err)
	}

	metaData, err := hex.DecodeString(valstr[130:])

	if err != nil {
		require.NoError(t, err)
	}

	println("metadata :", string(metaData))
}

func TestTransferFrom(t *testing.T) {
	return // ignore

	value := hex.EncodeToString(big.NewInt(3).Bytes())

	data := TransferFrom(key.Address, "0xcecbe670c11d4d28678955f23e0d2d708d79c893", value)

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
