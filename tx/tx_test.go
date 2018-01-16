package tx

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dynamicgo/config"
	"github.com/inwecrypto/ethgo"
	"github.com/inwecrypto/ethgo/erc20"
	"github.com/inwecrypto/ethgo/keystore"
	"github.com/inwecrypto/ethgo/rpc"
	"github.com/stretchr/testify/assert"
)

var key *keystore.Key

func init() {
	rawdata, err := ioutil.ReadFile("../../conf/keystore.json")

	if err != nil {
		panic(err)
	}

	key, err = keystore.ReadKeyStore(rawdata, "test")

	if err != nil {
		panic(err)
	}
}

var client *rpc.Client

func init() {
	cnf, _ := config.NewFromFile("../../conf/test.json")
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

	tx := NewTx(nonce, "0x96ae993fe6ac1786478d3d0b0eff780bff038276", nil, gasPrice, gasLimits, codes)

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

	tranferValue := ethgo.NewValue(big.NewFloat(0.1), ethgo.Ether)

	gasLimits := big.NewInt(21000)

	gasPrice := ethgo.NewValue(big.NewFloat(20), ethgo.Shannon)

	tx := NewTx(nonce, key.Address, tranferValue, gasPrice, gasLimits, nil)

	require.NoError(t, tx.Sign(key.PrivateKey))

	rawtx, err := tx.Encode()

	require.NoError(t, err)

	id, err := client.SendRawTransaction(rawtx)

	require.NoError(t, err)

	println(id)

}
