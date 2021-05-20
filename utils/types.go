package utils

import (
	"fmt"
)

//###################################################################################
// Constraint Value
//###################################################################################

type (
	Hash string // TODO [32]byte
	Name Hash
)

//###################################################################################
// Transaction
//###################################################################################

type Tx struct {
	from Hash
	to   Hash
}

//###################################################################################
// Constraint Message Value
//###################################################################################

type Cv struct {
	Hash  Hash
	Value interface{}
}

func (cv Cv) String() string {
	return fmt.Sprintf("%s: %v", cv.Hash, cv.Value)
}

type CvSet map[Hash]*Cv

func (set CvSet) Clean() {
	for k := range set {
		delete(set, k)
	}
}

func (set CvSet) HasCv() bool {
	return len(set) > 0
}

func (set CvSet) AddCv(cv *Cv) {
	var _cv = *cv
	set[cv.Hash] = &_cv
}

func (set CvSet) Copy() CvSet {
	var _set = make(CvSet, len(set))
	if set.HasCv() {
		for k, v := range set {
			_set[k] = v
		}
	}
	return _set
}

func (set CvSet) String() string {
	var format = "CvSet: {"
	for _, _v := range set {
		v := _v.String()
		format = fmt.Sprintf(format+"\n\t%v", v)
	}
	return format + "\n}"
}

//###################################################################################
// Blockchain
//###################################################################################

type TxSet CvSet

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
