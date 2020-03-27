package poa

import (
	"github.com/ldcc/mini-cluster/consensus"
	"github.com/ldcc/mini-cluster/utils"
)

type PoA struct {
	Engine *consensus.Engine
}

func MakePoA(name utils.Name) *consensus.Engine {
	self := PoA{}
	self.Engine = consensus.MakeEngine(name, self)
	return self.Engine
}
