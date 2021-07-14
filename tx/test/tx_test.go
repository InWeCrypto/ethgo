package tx

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ximenyan/ethgo"
	"github.com/ximenyan/ethgo/erc20"
	"github.com/ximenyan/ethgo/keystore"
	"github.com/ximenyan/ethgo/rpc"
	"github.com/ximenyan/ethgo/tx"
)

var key *keystore.Key

func init() {
	var err error
	s,_ := hex.DecodeString("1")
	key, err = keystore.KeyFromPrivateKey(s)
	if err != nil {
		panic(err)
	}
}

var client *rpc.Client

func init() {
	client = rpc.NewClient("https://bsc-dataseed1.binance.org")
}

func TestTokenTransfer(t *testing.T) {

	println(key.Address)

	deciamls, err := client.GetTokenDecimals("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")

	require.NoError(t, err)

	println("deciamls :", deciamls.Int64())

	balance, err := client.GetTokenBalance("0x7673d098c65621721e1eb91fcc76b3bbb0198945", key.Address)

	require.NoError(t, err)

	println("balance :", ethgo.CustomerValue(balance, deciamls).String())

	transferValue := ethgo.FromCustomerValue(big.NewFloat(10000), deciamls)

	codes, err := erc20.Transfer("0xF1cC4e0412E63c23dA7e3F881B60DE17F341aF65", hex.EncodeToString(transferValue.Bytes()))

	require.NoError(t, err)

	println("erc20 transfer code :", hex.EncodeToString(codes))

	gasLimits := big.NewInt(61000)

	gasPrice := ethgo.NewValue(big.NewFloat(20), ethgo.Shannon)

	nonce, err := client.Nonce(key.Address)

	require.NoError(t, err)

	tx := tx.NewTx(nonce, "0x6bf8c045ac5cc022568545997db24c946794c0c2", nil, gasPrice, gasLimits, codes)

	require.NoError(t, tx.Sign(key.PrivateKey))

	rawtx, err := tx.Encode()

	require.NoError(t, err)

	id, err := client.SendRawTransaction(rawtx)

	require.NoError(t, err)

	println(id)

}

func TestSign(t *testing.T) {

	println("test address", key.Address)

	nonce, err := client.Nonce(key.Address)

	assert.NoError(t, err)

	balance, err := client.GetBalance(key.Address)

	assert.NoError(t, err)

	println("nonce:", nonce)

	println("balance", fmt.Sprintf("%.018f", balance.As(ethgo.Ether)))

	tranferValue := ethgo.NewValue(big.NewFloat(0.01), ethgo.Ether)

	gasLimits := big.NewInt(90000)

	gasPrice := ethgo.NewValue(big.NewFloat(22), ethgo.Shannon)

	txdata := tx.NewTx(nonce, key.Address, tranferValue, gasPrice, gasLimits, nil)

	require.NoError(t, txdata.Sign(key.PrivateKey))

	rawtx, err := txdata.Encode()

	require.NoError(t, err)

	println(hex.EncodeToString(rawtx))

	// // rawtx, _ := hex.DecodeString("f8670b8083015f9094625e57af0057a4566255a2525303e68cdfe6606b872386f26fc10000801ba058684fd15bd356c67eb9cd24bb8ba20b866f6784d0665504e47ac3bd6f3baab6a069a6d78490c35ee13685abddd2d022ba966330c91a39bcf3e2abf41b60105d04")

	id, err := client.SendRawTransaction(rawtx)

	require.NoError(t, err)

	println(id)

}
