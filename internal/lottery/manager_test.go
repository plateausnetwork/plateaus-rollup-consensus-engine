package lottery

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/database"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/ipfs"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/nft"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/network"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager_SubscriberPeer(t *testing.T) {
	height := 1
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	latestBlock := LatestBlock{}
	latestBlock.Block.Header.Height = "1"

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(&latestBlock, nil)
	mockSubscriber.EXPECT().IsAvailable(gomock.Eq(&latestBlock)).Times(1).Return(true, nil)
	mockSubscriber.EXPECT().Subscribe(gomock.Eq(height), gomock.Eq(peer), gomock.Eq(networks)).Times(1).Return(nil)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: nil,
	}

	err := m.SubscribePeer(peer, networks)

	assert.NoError(t, err)
}

func TestManager_SubscriberPeer_GetLatestBlock_ZeroHeight(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	latestBlock := &LatestBlock{}
	latestBlock.Block.Header.Height = "0"
	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(latestBlock, nil)
	mockSubscriber.EXPECT().IsAvailable(gomock.Any()).Times(0)
	mockSubscriber.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: nil,
	}

	err := m.SubscribePeer(peer, networks)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "height must be greater than 0")
}

func TestManager_SubscriberPeer_GetLatestBlockError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(nil, errors.New("some error"))
	mockSubscriber.EXPECT().IsAvailable(gomock.Any()).Times(0)
	mockSubscriber.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: nil,
	}

	err := m.SubscribePeer(peer, networks)

	assert.Error(t, err)
}

func TestManager_SubscriberPeer_IsAvailableError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	latestBlock := LatestBlock{}
	latestBlock.Block.Header.Height = "1"

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(&latestBlock, nil)
	mockSubscriber.EXPECT().IsAvailable(gomock.Any()).Times(1).Return(false, errors.New("some error"))
	mockSubscriber.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: nil,
	}

	err := m.SubscribePeer(peer, networks)

	assert.Error(t, err)
}

func TestManager_SubscriberPeer_Subscribe(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	latestBlock := LatestBlock{}
	latestBlock.Block.Header.Height = "1"

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(&latestBlock, nil)
	mockSubscriber.EXPECT().IsAvailable(gomock.Any()).Times(1).Return(true, nil)
	mockSubscriber.EXPECT().Subscribe(gomock.Eq(latestBlock.GetHeight()), gomock.Eq(peer), gomock.Eq(networks)).Times(1).Return(errors.New("some error"))

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: nil,
	}

	err := m.SubscribePeer(peer, networks)

	assert.Error(t, err)
}

func TestManager_RegisterTx(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}
	txsMapped := &map[string]string{}
	ipfsURL := "ipfs://ipfs.io/test"
	latestBlock := &LatestBlock{}
	data := &database.Data{}
	nftValidation := &nft.LotteryValidation{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)
	mockDataRepository := database.NewMockDataRepository(ctrl)
	mockImageGenerator := nft.NewMockImageGenerator(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(latestBlock, nil)

	mockRegister.EXPECT().IsClosed().Times(1).Return(true, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Eq(peer)).Times(1).Return(networks[0], nil)
	mockRegister.EXPECT().PickWinner().Times(1)
	mockRegister.EXPECT().GenerateRoot(gomock.AssignableToTypeOf(&SubscribeBlocks{}), gomock.Eq(networks[0])).Times(1).Return(root, txsMapped, nil)

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(1).Return(false)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(1).Return(true)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Eq(ipfsURL), gomock.Any(), gomock.Any()).Times(1).Return(nil)

	mockImageGenerator.EXPECT().Generate(gomock.Eq(root.String())).Times(1).Return(nftValidation, nil)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, []nft.ImageGenerator{mockImageGenerator})

	mockDataRepository.EXPECT().Get().Times(1).Return(data, nil)

	mockClientIPFS := ipfs.NewMockClient(ctrl)
	mockClientIPFS.EXPECT().Put(gomock.Eq(context.TODO()), gomock.Eq(txsMapped)).Times(1).Return(ipfsURL, nil)
	serviceIPFS := ipfs.NewService(mockClientIPFS)

	m := Manager{
		s:    mockSubscriber,
		r:    mockRegister,
		dn:   delegator,
		dr:   mockDataRepository,
		ipfs: serviceIPFS,
	}

	err := m.RegisterTx(peer, networks)

	assert.NoError(t, err)
}

func TestManager_RegisterTx_LotteryNotClosed(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}
	latestBlock := LatestBlock{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)

	mockRegister.EXPECT().IsClosed().Times(1).Return(false, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Eq(peer)).Times(0)
	mockRegister.EXPECT().GenerateRoot(gomock.Eq(latestBlock), gomock.Eq(networks[0])).Times(0)

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, nil)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: delegator,
	}

	err := m.RegisterTx(peer, networks)

	assert.NoError(t, err)
}

func TestManager_RegisterTx_GetLotteryWinnersError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)

	mockRegister.EXPECT().IsClosed().Times(1).Return(true, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Eq(peer)).Times(1).Return("", errors.New("some error"))
	mockRegister.EXPECT().GenerateRoot(gomock.Any(), gomock.Eq(networks[0])).Times(0)
	mockRegister.EXPECT().PickWinner().Times(1).Return(nil)

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, nil)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: delegator,
	}

	err := m.RegisterTx(peer, networks)

	assert.Error(t, err)
}

func TestManager_RegisterTx_RegisterError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}
	latestBlock := &LatestBlock{}
	data := &database.Data{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)
	mockDataRepository := database.NewMockDataRepository(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(latestBlock, nil)

	mockRegister.EXPECT().IsClosed().Times(1).Return(true, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Eq(peer)).Times(1).Return(networks[0], nil)
	mockRegister.EXPECT().GenerateRoot(gomock.Any(), gomock.Eq(networks[0])).Times(1).Return(nil, nil, errors.New("some error"))
	mockRegister.EXPECT().PickWinner().Times(1).Return(nil)

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, nil)

	mockDataRepository.EXPECT().Get().Times(1).Return(data, nil)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: delegator,
		dr: mockDataRepository,
	}

	err := m.RegisterTx(peer, networks)

	assert.Error(t, err)
}

func TestManager_RegisterTx_MintNFTError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}
	txsMapped := &map[string]string{}
	ipfsURL := "ipfs://ipfs.io/test"
	latestBlock := &LatestBlock{}
	data := &database.Data{}
	nftValidation := &nft.LotteryValidation{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)
	mockDataRepository := database.NewMockDataRepository(ctrl)
	mockImageGenerator := nft.NewMockImageGenerator(ctrl)

	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(latestBlock, nil)

	mockRegister.EXPECT().IsClosed().Times(1).Return(true, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Eq(peer)).Times(1).Return(networks[0], nil)
	mockRegister.EXPECT().GenerateRoot(gomock.Any(), gomock.Eq(networks[0])).Times(1).Return(root, txsMapped, nil)
	mockRegister.EXPECT().PickWinner().Times(1).Return(nil)

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(1).Return(true)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Eq(ipfsURL), gomock.Any(), gomock.Any()).Times(1).Return(errors.New("some error"))
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	mockImageGenerator.EXPECT().Generate(gomock.Eq(root.String())).Times(1).Return(nftValidation, nil)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, []nft.ImageGenerator{mockImageGenerator})

	mockDataRepository.EXPECT().Get().Times(1).Return(data, nil)

	mockIPFSClient := ipfs.NewMockClient(ctrl)
	mockIPFSClient.EXPECT().Put(gomock.Eq(context.TODO()), gomock.Eq(txsMapped)).Times(1).Return(ipfsURL, nil)

	ipfsService := ipfs.NewService(mockIPFSClient)

	m := Manager{
		s:    mockSubscriber,
		r:    mockRegister,
		dn:   delegator,
		dr:   mockDataRepository,
		ipfs: ipfsService,
	}

	err := m.RegisterTx(peer, networks)

	assert.Error(t, err)
}

func TestManager_RegisterTx_PickWinnerError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)

	mockRegister.EXPECT().IsClosed().Times(1).Return(true, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Any()).Times(0)
	mockRegister.EXPECT().GenerateRoot(gomock.Any(), gomock.Eq(networks[0])).Times(0)
	mockRegister.EXPECT().PickWinner().Times(1).Return(errors.New("some error"))

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, nil)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: delegator,
	}

	err := m.RegisterTx(peer, networks)

	assert.Error(t, err)
}

func TestManager_RegisterTx_GetLatestBlockError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)

	mockRegister.EXPECT().IsClosed().Times(1).Return(true, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Any()).Times(1).Return(networks[0], nil)
	mockRegister.EXPECT().GenerateRoot(gomock.Any(), gomock.Eq(networks[0])).Times(0)
	mockRegister.EXPECT().PickWinner().Times(1).Return(nil)
	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(nil, errors.New("some error"))

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, nil)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: delegator,
	}

	err := m.RegisterTx(peer, networks)

	assert.Error(t, err)
}

func TestManager_RegisterTx_GetDataError(t *testing.T) {
	peer := "peer"
	networks := []string{"net_1", "net_2"}
	root := hash.Hash{}

	ctrl := gomock.NewController(t)

	mockSubscriber := NewMockSubscriber(ctrl)
	mockRegister := NewMockRegister(ctrl)
	mockDelegatedNet1 := network.NewMockDelegated(ctrl)
	mockDelegatedNet2 := network.NewMockDelegated(ctrl)
	mockDataRepository := database.NewMockDataRepository(ctrl)

	mockRegister.EXPECT().IsClosed().Times(1).Return(true, nil)
	mockRegister.EXPECT().GetLotteryWinners(gomock.Any()).Times(1).Return(networks[0], nil)
	mockRegister.EXPECT().GenerateRoot(gomock.Any(), gomock.Eq(networks[0])).Times(0)
	mockRegister.EXPECT().PickWinner().Times(1).Return(nil)
	mockSubscriber.EXPECT().GetLatestBlock().Times(1).Return(nil, nil)
	mockDataRepository.EXPECT().Get().Times(1).Return(nil, errors.New("some error"))

	mockDelegatedNet1.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet1.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockDelegatedNet2.EXPECT().Supports(gomock.Eq(networks[0])).Times(0)
	mockDelegatedNet2.EXPECT().MintNFT(gomock.Eq(root), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	delegator := network.NewDelegator([]network.Delegated{mockDelegatedNet1, mockDelegatedNet2}, nil)

	m := Manager{
		s:  mockSubscriber,
		r:  mockRegister,
		dn: delegator,
		dr: mockDataRepository,
	}

	err := m.RegisterTx(peer, networks)

	assert.Error(t, err)
}
