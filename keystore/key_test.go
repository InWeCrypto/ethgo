package keystore

import (
	"testing"
)

func TestN(t *testing.T) {
	println(LightScryptN, StandardScryptN)
}

// func TestKeyGen(t *testing.T) {
// 	key, err := NewKey()

// 	assert.NoError(t, err)

// 	println(key.Address)

// 	ks, err := WriteLightScryptKeyStore(key, "test")

// 	assert.NoError(t, err)

// 	key2, err := ReadKeyStore(ks, "test")

// 	assert.NoError(t, err)

// 	assert.Equal(t, key2.Address, key.Address)

// 	println(string(ks))

// 	privateKeyBytes := key.ToBytes()

// 	dic, _ := bip39.GetDict("en_US")

// 	data, _ := bip39.NewMnemonic(privateKeyBytes, dic)

// 	println(len(privateKeyBytes), len(strings.Split(data, " ")))

// 	println(data)

// 	data2, err := bip39.MnemonicToByteArray(data, dic)

// 	data2 = data2[1 : len(data2)-1]

// 	assert.NoError(t, err)

// 	assert.Equal(t, privateKeyBytes, data2)

// 	key3, err := KeyFromPrivateKey(data2)

// 	assert.NoError(t, err)

// 	assert.Equal(t, key.Address, key3.Address)
// }

// var key *Key

// func init() {
// 	rawdata, err := ioutil.ReadFile("../../conf/keystore.json")

// 	if err != nil {
// 		panic(err)
// 	}

// 	key, err = ReadKeyStore(rawdata, "test")

// 	if err != nil {
// 		panic(err)
// 	}
// }
