package tx

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/dynamicgo/config"
	"github.com/inwecrypto/ethgo"
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

func TestSign(t *testing.T) {

	println("test address", key.Address)

	nonce, err := client.Nonce(key.Address)

	assert.NoError(t, err)

	balance, err := client.GetBalance(key.Address)

	assert.NoError(t, err)

	println("nonce:", nonce)

	println("balance", fmt.Sprintf("%.018f", balance.As(ethgo.Ether)))

	tranferValue := ethgo.NewValue(big.NewFloat(0.1), ethgo.Ether)
	gasLimits := big.NewInt(200)
	gasPrice := ethgo.NewValue(big.NewFloat(20), ethgo.Wei)

	NewTx(nonce, key.Address, tranferValue, gasPrice, gasLimits, nil)

}
