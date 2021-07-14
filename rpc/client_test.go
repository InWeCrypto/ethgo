package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ximenyan/ethgo"
	"github.com/ximenyan/ethgo/atkcc"
	"github.com/ximenyan/ethgo/contract"
	"github.com/ximenyan/ethgo/erc20"
	"github.com/ximenyan/ethgo/keystore"
	"github.com/ximenyan/ethgo/tx"
	"github.com/ximenyan/ethgo/uniswap"
	"math/big"
	"testing"
	"time"

	"github.com/dynamicgo/config"
	"github.com/stretchr/testify/assert"
)

var cnf *config.Config
var client *Client
var TO_ADDRESS = "0xF1cC4e0412E63c23dA7e3F881B60DE17F341aF65"

var key *keystore.Key
var minOutAmount *big.Int
func init() {
	s,_ := hex.DecodeString("1")
	minOutAmount,_ = big.NewInt(0).SetString("700000000000000",10)
	key, _ = keystore.KeyFromPrivateKey(s)
	cnf, _ = config.NewFromFile("../../conf/test.json")
	client = NewClient("https://bsc-dataseed1.binance.org")
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

func ATKETH(_pair string)  {
	nonce,err := client.Nonce(TO_ADDRESS)
	codes := atkcc.GetATKETH(_pair)
	if err != nil{
		return
	}
	gasLimits := big.NewInt(524052)
	gasPrice := ethgo.NewValue(big.NewFloat(5.15), ethgo.Shannon)

	tx := tx.NewTx(nonce, atkcc.BSC_ATK_ADDRESS, nil, gasPrice, gasLimits, codes)
	tx.Sign(key.PrivateKey)

	rawtx, err := tx.Encode()
	id, err := client.SendRawTransaction(rawtx)
	fmt.Println("txid::::",id)
}

func getOutEth(_pair string) {
	data := atkcc.GetCanOutETH(_pair)
	valstr, _ := client.Call(&CallSite{
		To:   atkcc.BSC_ATK_ADDRESS,
		Data: data,
	})
	amount1,_ := ReadBigint(valstr[0:66])
	amount2,_ := ReadBigint(valstr[66:130])
	rev,_ := ReadBigint(valstr[130:])
	//if rev.Int64() != 0{
	//	amount1,amount2 = amount2, amount1
	//}
	if amount2.Cmp(minOutAmount) == 1{
		fmt.Println("getOutEth:",amount1.String(), amount2.String(), rev.String())
		ATKETH(_pair)
		//panic("xxxxxxxxxxxxxxxxxxxxx")
	}
}
func getPair(token string) (string,error){
	data := uniswap.GetPair(token, contract.WBNB)
	valstr, err := client.Call(&CallSite{
		To:   contract.PANCAKEV2_FACTORY_ADDR,
		Data: data,
	})
	if err != nil {
		return "",err
	} else{
		return valstr[26:], err
	}
	//
	//return ReadBigint(valstr)
	//var data string
	//
	//err = client.call("eth_getBalance", &data, address, "latest")
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//val, err := ReadBigint(data)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return (*ethgo.Value)(val), nil
}
func TestGetTransactionPool(t *testing.T) {
	for {
		txPool, err := client.GetTransactionPool()
		if err == nil{
			for _,v := range txPool.Pending{
				for _,_t := range v {
					if len(_t.Input) < 10{
						continue
					}
					_func := _t.Input[2:10]
					//_data := _t.Input[8:]
					if _func == erc20.TransferID{
						if _,ok := contract.BscNowPassed[_t.To]; ok{
							continue
						}else{
							pair,err := getPair(_t.To)
							if err == nil{
								if pair == "0000000000000000000000000000000000000000"{
									continue
								}
								t.Log("token txhash:", _t.Hash)
								getOutEth("0x"+pair)
							}
							//amountBytes,_ := hex.DecodeString(_data[len(_data)-64:])
							//t.Log(_t.Hash, _t.To ,_data[len(_data)-64:])
							//amount := big.NewInt(0).SetBytes(amountBytes)
							//t.Log("amount:",amount.String())
						}
					}
				}
			}
		}

		//time.Sleep(time.Millisecond * 500)
		//assert.NoError(t, err)
		//printResult(txPool.Pending)
	}
}

func TestGetBlockPending(t *testing.T) {
	for {
		blocknumber, err := client.GetBlockPending()
		assert.NoError(t, err)
		printResult(blocknumber.Number)
		time.Sleep(time.Second * 1)
	}
}

func TestGetBlockByNumber(t *testing.T) {

	blocknumber, err := client.GetBlockByNumber(9113024)

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
