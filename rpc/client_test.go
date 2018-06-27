package rpc

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dynamicgo/config"
	"github.com/stretchr/testify/assert"
)

var cnf *config.Config
var client *Client

func init() {
	cnf, _ = config.NewFromFile("../../conf/test.json")
	client = NewClient(cnf.GetString("ethmainnet", "http://xxxxxxx:8545"))
}

func TestBalance(t *testing.T) {

	accoutState, err := client.GetBalance("0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98")

	assert.NoError(t, err)

	printResult(accoutState)
}

func TestBlockNumber(t *testing.T) {

	blocknumber, err := client.BlockNumber()

	assert.NoError(t, err)

	printResult(blocknumber)
}

func TestBlocksPerSecond(t *testing.T) {

	blocknumber, err := client.BlockPerSecond()

	assert.NoError(t, err)

	printResult(blocknumber)
}

func TestGetBlockByNumber(t *testing.T) {

	blocknumber, err := client.GetBlockByNumber(5070477)

	assert.NoError(t, err)

	printResult(blocknumber)
}

func TestGetTransactionByHash(t *testing.T) {

	tx, err := client.GetTransactionByHash("0x525272f810bdb526e690c92886665bfdf41dab7f4626a77615cee32c3a66d93e")

	assert.NoError(t, err)

	printResult(tx)
}

func TestGetTransactionReceipt(t *testing.T) {

	tx, err := client.GetTransactionReceipt("0x73098500f6dcb8a42a7b7b56f095e4b17833b969c9bed25693381c6035c186ae")

	assert.NoError(t, err)

	printResult(tx)
}

func printResult(result interface{}) {

	data, _ := json.MarshalIndent(result, "", "\t")

	fmt.Println(string(data))
}
