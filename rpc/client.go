package rpc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/dynamicgo/slf4go"
	"github.com/ybbus/jsonrpc"
)

// Client neo jsonrpc 2.0 client
type Client struct {
	jsonrpcclient *jsonrpc.RPCClient
	slf4go.Logger
}

// NewClient create new neo client
func NewClient(url string) *Client {
	return &Client{
		jsonrpcclient: jsonrpc.NewRPCClient(url),
		Logger:        slf4go.Get("geth-rpc-client"),
	}
}

func (client *Client) call(method string, result interface{}, args ...interface{}) error {

	var buff bytes.Buffer

	buff.WriteString(fmt.Sprintf("jsonrpc call: %s\n", method))
	buff.WriteString(fmt.Sprintf("\tresult: %v\n", reflect.TypeOf(result)))

	for i, arg := range args {
		buff.WriteString(fmt.Sprintf("\targ(%d): %v\n", i, arg))
	}

	client.Debug(buff.String())

	response, err := client.jsonrpcclient.Call(method, args...)

	if err != nil {
		return err
	}

	if response.Error != nil {
		return fmt.Errorf("rpc error : %d %s %v", response.Error.Code, response.Error.Message, response.Error.Data)
	}

	buff.Reset()

	responsedata, _ := json.Marshal(response)

	buff.WriteString(fmt.Sprintf("jsonrpc call: %s\n", method))
	buff.WriteString(fmt.Sprintf("\tresult: %s\n", responsedata))

	client.Debug(buff.String())

	return response.GetObject(result)
}

// GetBalance get balance of eth address
func (client *Client) GetBalance(address string) (value *big.Int, err error) {

	var data string

	err = client.call("eth_getBalance", &data, address, "latest")

	if err != nil {
		return nil, err
	}

	return readBigint(data)
}

// BlockNumber get geth last block number
func (client *Client) BlockNumber() (uint64, error) {

	var data string

	err := client.call("eth_blockNumber", &data, "latest")

	if err != nil {
		return 0, err
	}

	val, err := readBigint(data)

	if err != nil {
		return 0, err
	}

	return val.Uint64(), nil
}

// BlockPerSecond get geth last block number
func (client *Client) BlockPerSecond() (val float64, err error) {

	err = client.call("blockPerSecond", &val)

	return
}

// Call eth call
func (client *Client) Call(callsite *CallSite) (val float64, err error) {

	err = client.call("eth_call", callsite, "latest", &val)

	return
}

// GetBlockByNumber get geth last block number
func (client *Client) GetBlockByNumber(number uint64) (val *Block, err error) {

	err = client.call("eth_getBlockByNumber", &val, fmt.Sprintf("0x%x", number), true)

	return
}

// GetTransactionByHash get geth last block number
func (client *Client) GetTransactionByHash(tx string) (val *Transaction, err error) {

	err = client.call("eth_getTransactionByHash", &val, tx)

	return
}

func readBigint(source string) (*big.Int, error) {
	value := big.NewInt(0)

	if source == "0x0" {
		return value, nil
	}

	if strings.HasPrefix(source, "0x") {
		source = source[2:]
	}

	data, err := hex.DecodeString(source)

	if err != nil {
		return nil, err
	}

	return value.SetBytes(data), nil
}
