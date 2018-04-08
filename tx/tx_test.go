package tx

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/inwecrypto/ethgo"
	"github.com/inwecrypto/ethgo/erc20"
	"github.com/inwecrypto/ethgo/keystore"
	"github.com/inwecrypto/ethgo/rpc"
)

var key *keystore.Key
var client *rpc.Client

var contractAddress = "0x4a60d261b55e10df8486f587929115238633d205"

//var contractAddress = "0x73c09b05b96688cfdf50aeae07e62509d84cd249"

func init() {
	pk := "fd300067b6adfaf826855ef688a73edd9ce30031a1d8245db2ca6bdb1491e080"

	bs, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}

	key, err = keystore.KeyFromPrivateKey(bs)

	//	rawdata := `{"address":"4a2d1208038f73cb634966c5b86f59153c003dfc","crypto":{"cipher":"aes-128-ctr","ciphertext":"5a50cfd7f3aabe0e9eab50a9a75ea346b072e24f4b29ffd598b4622fa354dfb8","cipherparams":{"iv":"12320e9ec5b3f47e200574666f257f5d"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bb12a9dcbc5f12870751e34382f73fbbe4a8c668015d10027b00529a121de991"},"mac":"ea066c8ac16bc430cbdff62b56a3a0b380a44c0a938b42c7a523dab17d805a5a"},"id":"ea748215-5e5e-468f-82f0-4074577f65c1","version":3}`
	//	var err error

	//	key, err = keystore.ReadKeyStore([]byte(rawdata), "123456")

	if err != nil {
		panic(err)
	}

	client = rpc.NewClient("http://127.0.0.1:8545")
}

func TestTokenTransfer(t *testing.T) {

	println(key.Address)

	deciamls := big.NewInt(8)

	balance, err := client.GetTokenBalance(contractAddress, key.Address)

	require.NoError(t, err)

	println("balance :", ethgo.CustomerValue(balance, deciamls).String())

	transferValue := ethgo.FromCustomerValue(big.NewFloat(0.2), deciamls)

	codes, err := erc20.TransferFrom("0x22087a6f5bbd4873b94234c5dba8a1808e3f2e9f", "0x3e53431a636a23a15d8f5c7c7f0de4fa2c6234ff", hex.EncodeToString(transferValue.Bytes()))

	require.NoError(t, err)

	gasLimits := big.NewInt(61000)

	gasPrice := ethgo.NewValue(big.NewFloat(20), ethgo.Shannon)

	amount := ethgo.NewValue(big.NewFloat(1000000), ethgo.Wei)

	nonce, err := client.Nonce(key.Address)

	require.NoError(t, err)

	tx := NewTx(nonce, contractAddress, amount, gasPrice, gasLimits, codes)

	require.NoError(t, tx.Sign(key.PrivateKey))

	rawtx, err := tx.Encode()

	require.NoError(t, err)

	id, err := client.SendRawTransaction(rawtx)

	require.NoError(t, err)

	println("txId: ", id)

}
