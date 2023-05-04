package lottery

//go:generate mockgen -source=$GOFILE -destination=./register_mock.go -package=$GOPACKAGE

import (
	"encoding/json"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/lottery/plateaus"
	"log"
	"sort"
)

type Register interface {
	GetLotteryWinners(peer string) (string, error)
	IsClosed() (bool, error)
	PickWinner() error
	GenerateRoot(subscribeBlocks *SubscribeBlocks) (hash.Hash, *map[string]string, error)
}

type RegisterService struct {
	http plateaus.HTTPClient
	rpc  plateaus.RPCClient
	hg   hash.Generator
}

// Verifying interface Compliance
var _ Register = (*RegisterService)(nil)

// NewRegisterService returns a new LotterySubscriber lottery.RegisterService
func NewRegisterService(c plateaus.HTTPClient, r plateaus.RPCClient, hg hash.Generator) Register {
	return &RegisterService{
		http: c,
		hg:   hg,
		rpc:  r,
	}
}

func (r RegisterService) GenerateRoot(subscribeBlocks *SubscribeBlocks) (hash.Hash, *map[string]string, error) {
	mtxs, err := r.getTransactions(subscribeBlocks.GetCurrentMinHeight(), subscribeBlocks.GetCurrentMaxHeight(), 0)

	if err != nil {
		return nil, nil, err
	}

	txs := sortTxs(mtxs)

	root, err := r.hg.GenerateByCollection(txs.ArrString())

	if err != nil {
		return nil, nil, err
	}

	mapHash, err := r.hg.GenerateByMap(txs.MapToHashAndRaw())

	return root, mapHash, err
}

func (r RegisterService) getTransactions(minHeight, maxHeight, page int) (*map[int]Txs, error) {
	limit := 1000
	offset := limit*page - 1

	if page == 0 {
		offset = 0
	}

	btxs, err := r.http.GetTransactions(minHeight, maxHeight, offset, limit)

	if err != nil {
		return nil, err
	}

	var txsResponse TxsResponse

	if err := json.Unmarshal(btxs, &txsResponse); err != nil {
		log.Printf("json.Unmarshal GetTransactions response: %s", err)
		return nil, err
	}

	pageTotal := txsResponse.Pagination.Total - (limit * page)

	if pageTotal > 0 && pageTotal > limit {
		// TODO: analisar a possibilidade de enviar requisições paralelas. O melhor caminho é a comunicação gRPC
		ret, err := r.getTransactions(minHeight, maxHeight, page+1)

		if err != nil {
			return nil, err
		}

		(*ret)[page] = txsResponse.Txs

		return ret, err
	}

	ret := map[int]Txs{
		page: txsResponse.Txs,
	}

	return &ret, nil
}

func sortTxs(mtxs *map[int]Txs) *Txs {
	keys := make([]int, 0, len(*mtxs))
	var txs Txs

	for k := range *mtxs {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		txs = append(txs, (*mtxs)[k]...)
	}

	return &txs
}

func (r RegisterService) IsClosed() (bool, error) {
	return r.rpc.IsClosed()
}

func (r RegisterService) GetLotteryWinners(peer string) (string, error) {
	res, err := r.rpc.GetLotteryWinners()

	return res, err
}

func (r RegisterService) PickWinner() error {
	ok, err := r.rpc.WasPicked()

	if err != nil {
		return err
	}

	if ok == true {
		log.Printf("winner was picked")
		return nil
	}

	return r.rpc.PickWinner()
}
