package pos

import (
	"github.com/ldcc/mini-cluster/consensus"
	"github.com/ldcc/mini-cluster/utils"
)

type PoS struct {
	Engine *consensus.Engine
}

func MakePoS(name utils.Name) *consensus.Engine {
	self := PoS{}
	self.Engine = consensus.MakeEngine(name, self)
	return self.Engine
}
