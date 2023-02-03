package database

const EachBlock = 1000

type Data struct {
	LastBlockSubscribed int `json:"last_block_subscribed"`
}

func (d Data) GetNextBlock() int {
	return d.LastBlockSubscribed + EachBlock
}
