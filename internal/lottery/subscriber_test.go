package lottery

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/database"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/plateaus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_GetLatestBlock(t *testing.T) {
	latestBlock := []byte(`{"block":{"header":{"chain_id":"plateaus_432-1","height":"1"}}}`)

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientHTTPMock := plateaus.NewMockHTTPClient(ctrl)

	clientHTTPMock.EXPECT().GetLatestBlock().Times(1).Return(latestBlock, nil)

	s := SubscribeService{
		dr:   repositoryMock,
		http: clientHTTPMock,
	}

	l, err := s.GetLatestBlock()

	assert.NoError(t, err)
	assert.Equal(t, l.Block.Header.Height, "1")
	assert.Equal(t, l.GetHeight(), 1)
}

func TestService_GetLatestBlock_UnmarshalLatestBlock_Error(t *testing.T) {
	latestBlock := []byte(`{"block":{"header":{"chain_id":"plateaus_432-1","height":"1"`)

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientHTTPMock := plateaus.NewMockHTTPClient(ctrl)

	clientHTTPMock.EXPECT().GetLatestBlock().Times(1).Return(latestBlock, nil)

	s := SubscribeService{
		dr:   repositoryMock,
		http: clientHTTPMock,
	}

	_, err := s.GetLatestBlock()

	assert.Error(t, err)
}

func TestService_GetLatestBlock_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientHTTPMock := plateaus.NewMockHTTPClient(ctrl)

	clientHTTPMock.EXPECT().GetLatestBlock().Times(1).Return(nil, errors.New("some error"))

	s := SubscribeService{
		dr:   repositoryMock,
		http: clientHTTPMock,
	}

	_, err := s.GetLatestBlock()

	assert.Error(t, err)
}

func TestService_IsAvailable(t *testing.T) {
	latestBlock := &LatestBlock{}
	latestBlock.Block.Header.Height = "100000"
	data := database.Data{
		LastBlockSubscribed: 1,
	}

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientRPCMock := plateaus.NewMockRPCClient(ctrl)

	repositoryMock.EXPECT().Get().Times(1).Return(&data, nil)
	clientRPCMock.EXPECT().IsOpen(gomock.AssignableToTypeOf(time.Time{})).Times(1).Return(true, nil)

	s := SubscribeService{
		dr:  repositoryMock,
		rpc: clientRPCMock,
	}

	b, err := s.IsAvailable(latestBlock)

	assert.NoError(t, err)
	assert.True(t, b)
}

func TestService_IsAvailable_LatestBlocksIsNotANumber_Error(t *testing.T) {
	latestBlock := &LatestBlock{}
	latestBlock.Block.Header.Height = "BBB"

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientHTTPMock := plateaus.NewMockHTTPClient(ctrl)

	repositoryMock.EXPECT().Get().Times(0)

	s := SubscribeService{
		dr:   repositoryMock,
		http: clientHTTPMock,
	}

	b, err := s.IsAvailable(latestBlock)

	assert.Error(t, err)
	assert.False(t, b)
}

func TestService_IsAvailable_HeightIsLessThanLatestBlocksSubscribed(t *testing.T) {
	latestBlock := &LatestBlock{}
	latestBlock.Block.Header.Height = "1"
	data := database.Data{
		LastBlockSubscribed: 1000,
	}

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientHTTPMock := plateaus.NewMockHTTPClient(ctrl)

	repositoryMock.EXPECT().Get().Times(1).Return(&data, nil)

	s := SubscribeService{
		dr:   repositoryMock,
		http: clientHTTPMock,
	}

	b, err := s.IsAvailable(latestBlock)

	assert.NoError(t, err)
	assert.False(t, b)
}

func TestService_IsAvailable_HeightIsLessThanLastBlockSubscribed(t *testing.T) {
	latestBlock := &LatestBlock{}
	latestBlock.Block.Header.Height = "1"
	data := database.Data{
		LastBlockSubscribed: 1000,
	}

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientHTTPMock := plateaus.NewMockHTTPClient(ctrl)

	repositoryMock.EXPECT().Get().Times(1).Return(&data, nil)

	s := SubscribeService{
		dr:   repositoryMock,
		http: clientHTTPMock,
	}

	b, err := s.IsAvailable(latestBlock)

	assert.NoError(t, err)
	assert.False(t, b)
}

func TestService_IsAvailable_LastBlockSubscribedIsGreaterNextMinHeight(t *testing.T) {
	latestBlock := &LatestBlock{}
	latestBlock.Block.Header.Height = "990"
	data := database.Data{
		LastBlockSubscribed: 200,
	}

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientHTTPMock := plateaus.NewMockHTTPClient(ctrl)

	repositoryMock.EXPECT().Get().Times(1).Return(&data, nil)

	s := SubscribeService{
		dr:   repositoryMock,
		http: clientHTTPMock,
	}

	b, err := s.IsAvailable(latestBlock)

	assert.NoError(t, err)
	assert.False(t, b)
}

func TestService_Subscribe(t *testing.T) {
	peer := "peer"
	networks := []string{
		"polygon", "arbitrum",
	}

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientRPCMock := plateaus.NewMockRPCClient(ctrl)

	clientRPCMock.EXPECT().Subscribe(gomock.Any()).Times(2).Return(nil)
	repositoryMock.EXPECT().Store(gomock.Any()).Times(1).Return(nil)

	s := SubscribeService{
		dr:  repositoryMock,
		rpc: clientRPCMock,
	}

	err := s.Subscribe(1, peer, networks)

	assert.NoError(t, err)
}

func TestService_Subscribe_RPCSubscribe_Error(t *testing.T) {
	peer := "peer"
	networks := []string{
		"polygon", "arbitrum",
	}

	ctrl := gomock.NewController(t)
	repositoryMock := database.NewMockDataRepository(ctrl)
	clientRPCMock := plateaus.NewMockRPCClient(ctrl)

	clientRPCMock.EXPECT().Subscribe(gomock.Eq(networks[0])).Times(1).Return(errors.New("some error"))
	repositoryMock.EXPECT().Store(gomock.Any()).Times(0)

	s := SubscribeService{
		dr:  repositoryMock,
		rpc: clientRPCMock,
	}

	err := s.Subscribe(1, peer, networks)

	assert.Error(t, err)
}
