package consensus

import (
	"github.com/ldcc/mini-cluster/utils"
)

//###################################################################################
// Consensus Engine Typeclass
//###################################################################################
type engineI interface {
	// TODO design a suit of `consensus engine interfaces`
}
type Engine struct {
	Name utils.Name
}

func MakeEngine(name utils.Name, ei engineI) *Engine {
	return &Engine{name}
}
