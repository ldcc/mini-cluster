package utils

import "fmt"

type Name Hash
type Hash string // [32]byte

/// Constraint Value
type Cv struct {
	Value interface{}
}

type Tx struct {
	Hash Hash
}

type TxSet map[Hash]*Tx

func (set TxSet) Clean() {
	for k := range set {
		delete(set, k)
	}
}

func (set TxSet) HasTx() bool {
	return len(set) > 0
}

func (set TxSet) AddTx(tx *Tx) {
	set[tx.Hash] = tx
}

func (set TxSet) Copy() TxSet {
	var txSet = make(TxSet)
	if set.HasTx() {
		for k, v := range set {
			txSet[k] = v
		}
	}
	return txSet
}

func (set TxSet) String() string {
	var format = "TxSet: {"
	for k, v := range set {
		format = fmt.Sprintf(format+"\n\t%v: %v,", k, *v)
	}
	return format + "\n}"
}

type CvList []*Cv

func (cvl *CvList) Clean() {
	tmplst := *cvl
	*cvl = tmplst[:0]
}

func (cvl *CvList) HasCv() bool {
	return len(*cvl) > 0
}

func (cvl *CvList) AddCv(cv *Cv) {
	*cvl = append(*cvl, cv)
}

func (cvl *CvList) Copy() CvList {
	var txList = make(CvList, len(*cvl))
	if cvl.HasCv() {
		for _, e := range *cvl {
			txList = append(txList, e)
		}
	}
	return txList
}

func (cvl CvList) String() string {
	var format = "CvList: {"
	for k, v := range cvl {
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
