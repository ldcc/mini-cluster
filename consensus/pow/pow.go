package pow

import (
	"github.com/ldcc/mini-cluster/consensus"
	"github.com/ldcc/mini-cluster/utils"
)

type PoW struct {
	Engine *consensus.Engine
}

func MakePoW(name utils.Name) *consensus.Engine {
	self := PoW{}
	self.Engine = consensus.MakeEngine(name, self)
	return self.Engine
}
