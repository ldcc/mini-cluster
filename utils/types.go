package utils

type Name Hash
type Hash string

type Block struct {
}

type Chain []*Block

type Tx struct {
}

type TxList []*Tx

func (self *TxList) Clean() {
	*self = nil
}

func (self *TxList) HasTx() bool {
	return len(*self) > 0
}

func (self *TxList) AddTx(tx Tx) {
	*self = append(*self, &tx)
}
