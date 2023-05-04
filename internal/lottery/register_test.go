package lottery

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/plateaus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister_Register(t *testing.T) {
	subscribedBlocks := &SubscribeBlocks{
		block:      1055,
		subscribed: 900,
		eachBlock:  1000,
	}

	txsBytes := []byte(`{
		"tx_responses": [{
			"txhash": "hash_1",
			"raw_log": "raw_log_1"
		}]
	}`)
	txsArrString := &[]string{"raw_log_1"}
	txsMapString := &map[string]string{"hash_1": "raw_log_1"}

	ctrl := gomock.NewController(t)

	mockClientHTTP := plateaus.NewMockHTTPClient(ctrl)
	mockClientRPC := plateaus.NewMockRPCClient(ctrl)
	mockHashGenerator := hash.NewMockGenerator(ctrl)

	mockClientHTTP.EXPECT().GetTransactions(gomock.Eq(0), gomock.Eq(999), gomock.Eq(0), gomock.Eq(1000)).Times(1).Return(txsBytes, nil)
	mockHashGenerator.EXPECT().GenerateByCollection(gomock.Eq(txsArrString)).Times(1).Return(hash.Hash{}, nil)
	mockHashGenerator.EXPECT().GenerateByMap(gomock.Eq(txsMapString)).Times(1).Return(txsMapString, nil)

	m := RegisterService{
		http: mockClientHTTP,
		rpc:  mockClientRPC,
		hg:   mockHashGenerator,
	}

	res, txs, err := m.GenerateRoot(subscribedBlocks)

	assert.IsType(t, hash.Hash{}, res)
	assert.IsType(t, &map[string]string{}, txs)
	assert.Equal(t, txsMapString, txs)
	assert.NoError(t, err)
}

func TestRegister_Register_GetTransactions_11_Times(t *testing.T) {
	subscribedBlocks := &SubscribeBlocks{
		block:      1055,
		subscribed: 900,
		eachBlock:  1000,
	}

	txsBytes := []byte(`
	{
		"tx_responses": [{
			"txhash": "hash_1",
			"raw_log": "raw_log_1"
		}],
		"pagination": {
			"next_key": null,
			"total": "11000"
		}
	}`)
	txsArrString := &[]string{
		"raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1", "raw_log_1",
	}
	txsMapString := &map[string]string{
		"hash_1": "raw_log_1",
	}

	ctrl := gomock.NewController(t)

	mockClientHTTP := plateaus.NewMockHTTPClient(ctrl)
	mockClientRPC := plateaus.NewMockRPCClient(ctrl)
	mockHashGenerator := hash.NewMockGenerator(ctrl)

	mockClientHTTP.EXPECT().GetTransactions(gomock.Eq(0), gomock.Eq(999), gomock.Any(), gomock.Eq(1000)).Times(11).Return(txsBytes, nil)
	mockHashGenerator.EXPECT().GenerateByCollection(gomock.Eq(txsArrString)).Times(1).Return(hash.Hash{}, nil)
	mockHashGenerator.EXPECT().GenerateByMap(gomock.Eq(txsMapString)).Times(1).Return(txsMapString, nil)

	m := RegisterService{
		http: mockClientHTTP,
		rpc:  mockClientRPC,
		hg:   mockHashGenerator,
	}

	res, tx, err := m.GenerateRoot(subscribedBlocks)

	assert.IsType(t, hash.Hash{}, res)
	assert.IsType(t, &map[string]string{}, tx)
	assert.NoError(t, err)
}

func TestRegister_PickWinner(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockClientRPC := plateaus.NewMockRPCClient(ctrl)
	mockClientRPC.EXPECT().WasPicked().Times(1).Return(false, nil)
	mockClientRPC.EXPECT().PickWinner().Times(1).Return(nil)

	m := RegisterService{
		http: nil,
		rpc:  mockClientRPC,
		hg:   nil,
	}

	err := m.PickWinner()

	assert.NoError(t, err)
}

func TestRegister_PickWinner_PickWinner_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockClientRPC := plateaus.NewMockRPCClient(ctrl)
	mockClientRPC.EXPECT().WasPicked().Times(1).Return(false, nil)
	mockClientRPC.EXPECT().PickWinner().Times(1).Return(errors.New("some error"))

	m := RegisterService{
		http: nil,
		rpc:  mockClientRPC,
		hg:   nil,
	}

	err := m.PickWinner()

	assert.Error(t, err)
}

func TestRegister_PickWinner_WasPicked_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockClientRPC := plateaus.NewMockRPCClient(ctrl)
	mockClientRPC.EXPECT().WasPicked().Times(1).Return(false, errors.New("some error"))
	mockClientRPC.EXPECT().PickWinner().Times(0)

	m := RegisterService{
		http: nil,
		rpc:  mockClientRPC,
		hg:   nil,
	}

	err := m.PickWinner()

	assert.Error(t, err)
}

func TestRegister_PickWinner_WinnerWasPicked(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockClientRPC := plateaus.NewMockRPCClient(ctrl)
	mockClientRPC.EXPECT().WasPicked().Times(1).Return(true, nil)
	mockClientRPC.EXPECT().PickWinner().Times(0)

	m := RegisterService{
		http: nil,
		rpc:  mockClientRPC,
		hg:   nil,
	}

	err := m.PickWinner()

	assert.NoError(t, err)
}

func TestRegister_PickWinner_WinnerWasPicked_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockClientRPC := plateaus.NewMockRPCClient(ctrl)
	mockClientRPC.EXPECT().WasPicked().Times(1).Return(false, errors.New("some error"))
	mockClientRPC.EXPECT().PickWinner().Times(0)

	m := RegisterService{
		http: nil,
		rpc:  mockClientRPC,
		hg:   nil,
	}

	err := m.PickWinner()

	assert.Error(t, err)
}
