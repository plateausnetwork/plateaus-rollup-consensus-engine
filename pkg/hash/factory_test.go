package merkletree

import (
	"encoding/hex"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery"
	"github.com/stretchr/testify/assert"
	"github.com/txaty/go-merkletree"
	"testing"
)

type TestStub struct {
	X string
}

//Serialize hashes the values of a DataBlock
func (t TestStub) Serialize() ([]byte, error) {
	return []byte(t.X), nil
}

func TestFactory_GenerateByCollection(t *testing.T) {
	txs := &[]string{"raw-log-1", "raw-log-2", "raw-log-3", "raw-log-4", "raw-log-5"}

	f := NewFactory()

	root, err := f.GenerateByCollection(txs)

	assert.NoError(t, err)
	assert.Equal(t, "b5ea4de4dec946f479d32bc2993cf50eded2951fa301597141e2c31c5fa4422f", root.String())
	assert.IsType(t, hash.Hash{}, root)
}

func TestFactory_GenerateByCollection_EmptyList_Error(t *testing.T) {
	txs := &[]string{}

	f := NewFactory()

	_, err := f.GenerateByCollection(txs)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "the number of data blocks must be greater than 1")
}

func TestFactory_GenerateByMap(t *testing.T) {
	txs := &map[string]string{
		"hash-1": "raw-log-1",
		"hash-2": "raw-log-2",
		"hash-3": "raw-log-3",
	}

	f := NewFactory()

	txsMapped, err := f.GenerateByMap(txs)

	assert.NoError(t, err)
	assert.Equal(t, "84ba122a2da788fbed2c52976c6062ad93d8e0eb4917c87092c7475a35b8cd10", (*txsMapped)["hash-1"])
	assert.Equal(t, "9aecd85187c5d8fffe5363bb352e9c3966d8267241a3df1bf7e73f269e6b2a9d", (*txsMapped)["hash-2"])
	assert.Equal(t, "6c6eb796a127bcd8c1a6cff3eee4b862bc061a3b9bcaa6b3223aa9c03667035c", (*txsMapped)["hash-3"])
	assert.IsType(t, &map[string]string{}, txsMapped)
}

func TestFactory_MatchingValues(t *testing.T) {
	txs := &lottery.Txs{
		lottery.Tx{
			Txhash: "hash-1",
			RawLog: "raw-log-1",
		},
		lottery.Tx{
			Txhash: "hash-2",
			RawLog: "raw-log-2",
		},
		lottery.Tx{
			Txhash: "hash-3",
			RawLog: "raw-log-3",
		}, lottery.Tx{
			Txhash: "hash-4",
			RawLog: "raw-log-4",
		},
		lottery.Tx{
			Txhash: "hash-5",
			RawLog: "raw-log-5",
		},
	}
	var txsKeys []string

	for _, v := range *txs {
		txsKeys = append(txsKeys, v.Txhash)
	}

	f := NewFactory()

	txsArr := txs.ArrString()
	root, _ := f.GenerateByCollection(txsArr)
	val := txs.MapToHashAndRaw()
	txsMapped, _ := f.GenerateByMap(val)
	root2 := generateRootManual(txsKeys, *txsMapped)

	assert.Equal(t, root, root2)
}

func TestFactory_OneItemOnList(t *testing.T) {
	txs := &lottery.Txs{
		lottery.Tx{
			Txhash: "hash-1",
			RawLog: "the message to hash here",
		},
	}
	var txsKeys []string

	for _, v := range *txs {
		txsKeys = append(txsKeys, v.Txhash)
	}

	f := NewFactory()

	root, _ := f.GenerateByCollection(txs.ArrString())
	val := txs.MapToHashAndRaw()
	txsMapped, _ := f.GenerateByMap(val)
	root2 := generateRootManual(txsKeys, *txsMapped)

	assert.Equal(t, root, root2)
}

func TestHash(t *testing.T) {
	txs := &lottery.Txs{
		lottery.Tx{
			Txhash: "hash-1",
			RawLog: "the message to hash here",
		},
		lottery.Tx{
			Txhash: "hash-2",
			RawLog: "the message to hash here",
		},
	}
	var txsKeys []string

	for _, v := range *txs {
		txsKeys = append(txsKeys, v.Txhash)
	}

	f := NewFactory()

	root, _ := f.GenerateByCollection(txs.ArrString())
	val := txs.MapToHashAndRaw()
	txsMapped, _ := f.GenerateByMap(val)
	root2 := generateRootManual(txsKeys, *txsMapped)

	assert.Equal(t, root, root2)
}

func generateRootManual(keys []string, elements map[string]string) hash.Hash {
	var list []merkletree.DataBlock

	for _, key := range keys {
		v := elements[key]
		// TODO: decouple TestContent
		list = append(list, TestStub{X: v})
	}

	if len(list) == 1 {
		list = append(list, TestContent{X: ""})
	}

	tree, _ := merkletree.New(&merkletree.Config{
		HashFunc: func(bytes []byte) ([]byte, error) {
			h, err := hex.DecodeString(string(bytes))

			if err != nil || string(bytes) == "" {
				return Sha256Func(bytes)
			}

			return h, err
		},
	}, list)

	return tree.Root
}

func TestFactory_Test(t *testing.T) {
	txs := &[]string{
		`[{"events":[{"type":"ethereum_tx","attributes":[{"key":"amount","value":"0"},{"key":"ethereumTxHash","value":"0x76ba8532a31727cdd1c9e5a62ade36b78cdd4f03199cd1526e4ebd431a6b41df"},{"key":"txIndex","value":"0"},{"key":"txGasUsed","value":"1500000"},{"key":"txHash","value":"6DAEBB7B6DD944A92C3C5B0EE7328D1D0E0BB85176A75BA82409F9AA131DF07C"},{"key":"recipient","value":"0x55352d22501abFDce6d8Cb60B47aeE8126C6694A"}]},{"type":"message","attributes":[{"key":"action","value":"/ethermint.evm.v1.MsgEthereumTx"},{"key":"module","value":"evm"},{"key":"sender","value":"0xEe0A356f16e895eE69Ea618C08aDd07b11dE1D9a"},{"key":"txType","value":"0"}]},{"type":"tx_log","attributes":null}]}]`,
	}

	f := NewFactory()

	root, err := f.GenerateByCollection(txs)

	assert.NoError(t, err)
	assert.Equal(t, "0e5db8cf65e22da8d600b5cc5b8ba612b8548ffb545b53e09575b46125eb0fa4", root.String())
	assert.IsType(t, hash.Hash{}, root)
}
