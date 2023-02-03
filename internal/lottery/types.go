package lottery

import (
	"fmt"
	"log"
	"time"
)

type SubscribeBlocks struct {
	block      int
	subscribed int
	eachBlock  int
}

func NewSubscribeBlocks(latest, lastSubscribed, eachBlock int) *SubscribeBlocks {
	return &SubscribeBlocks{
		block:      latest,
		subscribed: lastSubscribed,
		eachBlock:  eachBlock,
	}
}

func (s SubscribeBlocks) GetCurrentMinHeight() int {
	extraHeight := s.block % s.eachBlock

	return s.block - extraHeight - s.eachBlock
}

func (s SubscribeBlocks) GetCurrentMaxHeight() int {
	return s.GetCurrentMinHeight() + s.eachBlock - 1
}

func (s SubscribeBlocks) GetLastMaxHeight() int {
	return s.GetCurrentMinHeight() - 1
}

func (s SubscribeBlocks) GetHeightWaiting() int {
	extraHeight := s.subscribed % s.eachBlock

	return s.subscribed - extraHeight + s.eachBlock
}

type LatestBlock struct {
	Block struct {
		Header struct {
			Version struct {
				Block string `json:"block"`
			} `json:"version"`
			ChainID string    `json:"chain_id"`
			Height  string    `json:"height"`
			Time    time.Time `json:"time"`
		} `json:"header"`
	} `json:"block"`
}

func (lb LatestBlock) GetHeight() int {
	var height int
	if _, err := fmt.Sscan(lb.Block.Header.Height, &height); err != nil {
		log.Printf("could not convert LatestBlock.Height to int: %s", err)
		return 0
	}

	return height
}

type AutoGenerated struct {
	TxResponses []struct {
		Height    string `json:"height"`
		Txhash    string `json:"txhash"`
		Codespace string `json:"codespace"`
		Code      int    `json:"code"`
		Data      string `json:"data"`
		RawLog    string `json:"raw_log"`
		Logs      []struct {
			MsgIndex int    `json:"msg_index"`
			Log      string `json:"log"`
			Events   []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"logs"`
		Info      string `json:"info"`
		GasWanted string `json:"gas_wanted"`
		GasUsed   string `json:"gas_used"`
		Tx        struct {
			Type string `json:"@type"`
			Body struct {
				Messages []struct {
					Type string `json:"@type"`
					Data struct {
						Type     string `json:"@type"`
						Nonce    string `json:"nonce"`
						GasPrice string `json:"gas_price"`
						Gas      string `json:"gas"`
						To       string `json:"to"`
						Value    string `json:"value"`
						Data     string `json:"data"`
						V        string `json:"v"`
						R        string `json:"r"`
						S        string `json:"s"`
					} `json:"data"`
					Size int    `json:"size"`
					Hash string `json:"hash"`
					From string `json:"from"`
				} `json:"messages"`
				Memo             string `json:"memo"`
				TimeoutHeight    string `json:"timeout_height"`
				ExtensionOptions []struct {
					Type string `json:"@type"`
				} `json:"extension_options"`
				NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
			} `json:"body"`
			AuthInfo struct {
				SignerInfos []interface{} `json:"signer_infos"`
				Fee         struct {
					Amount   []interface{} `json:"amount"`
					GasLimit string        `json:"gas_limit"`
					Payer    string        `json:"payer"`
					Granter  string        `json:"granter"`
				} `json:"fee"`
			} `json:"auth_info"`
			Signatures []interface{} `json:"signatures"`
		} `json:"tx"`
		Timestamp time.Time `json:"timestamp"`
		Events    []struct {
			Type       string `json:"type"`
			Attributes []struct {
				Key   string      `json:"key"`
				Value interface{} `json:"value"`
				Index bool        `json:"index"`
			} `json:"attributes"`
		} `json:"events"`
	} `json:"tx_responses"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

type TxsResponse struct {
	Txs        []Tx       `json:"tx_responses"`
	Pagination Pagination `json:"pagination"`
}

type Tx struct {
	Height string `json:"height"`
	Txhash string `json:"txhash"`
	Data   string `json:"data"`
	RawLog string `json:"raw_log"`
	Logs   []struct {
		Events []struct {
			Type       string `json:"type"`
			Attributes []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"attributes"`
		} `json:"events"`
	} `json:"logs"`
	GasWanted string `json:"gas_wanted"`
	GasUsed   string `json:"gas_used"`
	Tx        struct {
		Type  string `json:"type"`
		Value struct {
			Msg []struct {
				Type  string `json:"type"`
				Value struct {
					FromAddress string `json:"from_address"`
					ToAddress   string `json:"to_address"`
					Amount      []struct {
						Denom  string `json:"denom"`
						Amount string `json:"amount"`
					} `json:"amount"`
				} `json:"value"`
			} `json:"msg"`
			Fee struct {
				Amount []interface{} `json:"amount"`
				Gas    string        `json:"gas"`
			} `json:"fee"`
			Signatures    []interface{} `json:"signatures"`
			Memo          string        `json:"memo"`
			TimeoutHeight string        `json:"timeout_height"`
		} `json:"value"`
	} `json:"tx"`
	Timestamp time.Time `json:"timestamp"`
	Events    []struct {
		Type       string `json:"type"`
		Attributes []struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
			Index bool        `json:"index"`
		} `json:"attributes"`
	} `json:"events"`
}

type Txs []Tx

func (t Txs) ArrString() *[]string {
	var arr []string

	for _, v := range t {
		arr = append(arr, v.RawLog)
	}

	return &arr
}

func (t Txs) MapToHashAndRaw() *map[string]string {
	res := map[string]string{}

	for _, tx := range t {
		res[tx.Txhash] = tx.RawLog
	}

	return &res
}

type Pagination struct {
	NextKey interface{} `json:"next_key"`
	Total   int         `json:"total,string"`
}