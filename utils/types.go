package utils

import "fmt"

type Name Hash
type Hash string // [32]byte

type Tx struct {
	Hash Hash
}

type TxSet map[Hash]*Tx

func (self TxSet) Clean() {
	for k := range self {
		delete(self, k)
	}
}

func (self TxSet) HasTx() bool {
	return len(self) > 0
}

func (self TxSet) AddTx(tx Tx) {
	self[tx.Hash] = &tx
}

func (self TxSet) Copy() TxSet {
	var txSet = make(TxSet)
	if self.HasTx() {
		for k, v := range self {
			txSet[k] = v
		}
	}
	return txSet
}

func (self TxSet) String() string {
	var format = "TxSet: {"
	for k, v := range self {
		format = fmt.Sprintf(format+"\n\t%v: %v,", k, *v)
	}
	return format + "\n}"
}

type TxList []*Tx

func (self *TxList) Clean() {
	tmplst := *self
	*self = tmplst[:0]
}

func (self *TxList) HasTx() bool {
	return len(*self) > 0
}

func (self *TxList) AddTx(tx Tx) {
	*self = append(*self, &tx)
}

func (self *TxList) Copy() TxList {
	var txList = make(TxList, len(*self))
	if self.HasTx() {
		for _, e := range *self {
			txList = append(txList, e)
		}
	}
	return txList
}

func (self TxList) String() string {
	var format = "TxList: {"
	for k, v := range self {
		format = fmt.Sprintf(format+"\n\t%No.n: %v,", k, *v)
	}
	return format + "\n}"
}

type Block struct {
	RootHash Hash
	PrvHash  Hash
	TxSet    TxSet
}

type Chain struct {
	NetworkId uint32
	RootHash  Hash
	blocks    []*Block
}

func GenesisChain(id uint32, name Name) *Chain {
	return &Chain{id, Hash(name), make([]*Block, 0)}
}
