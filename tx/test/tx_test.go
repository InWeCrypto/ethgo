package tx

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dynamicgo/config"
	"github.com/inwecrypto/ethgo"
	"github.com/inwecrypto/ethgo/erc20"
	"github.com/inwecrypto/ethgo/keystore"
	"github.com/inwecrypto/ethgo/rpc"
	"github.com/inwecrypto/ethgo/tx"
)

var key *keystore.Key

func init() {
	rawdata, err := ioutil.ReadFile("../../../conf/keystore.json")

	if err != nil {
		panic(err)
	}

	key, err = keystore.ReadKeyStore(rawdata, "Lalala123")

	if err != nil {
		panic(err)
	}
}

var client *rpc.Client

func init() {
	cnf, _ := config.NewFromFile("../../../conf/test.json")
	client = rpc.NewClient(cnf.GetString("ethtestnet", "http://xxxxxxx:8545"))
}

func TestTokenTransfer(t *testing.T) {

	println(key.Address)

	deciamls, err := client.GetTokenDecimals("0x96ae993fe6ac1786478d3d0b0eff780bff038276")

	require.NoError(t, err)

	println("deciamls :", deciamls.Int64())

	balance, err := client.GetTokenBalance("0x96ae993fe6ac1786478d3d0b0eff780bff038276", key.Address)

	require.NoError(t, err)

	println("balance :", ethgo.CustomerValue(balance, deciamls).String())

	transferValue := ethgo.FromCustomerValue(big.NewFloat(200.0), deciamls)

	codes, err := erc20.Transfer(key.Address, hex.EncodeToString(transferValue.Bytes()))

	require.NoError(t, err)

	println("erc20 transfer code :", hex.EncodeToString(codes))

	gasLimits := big.NewInt(61000)

	gasPrice := ethgo.NewValue(big.NewFloat(20), ethgo.Shannon)

	nonce, err := client.Nonce(key.Address)

	require.NoError(t, err)

	tx := tx.NewTx(nonce, "0x96ae993fe6ac1786478d3d0b0eff780bff038276", nil, gasPrice, gasLimits, codes)

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
