package consensus

import (
	"../utils"
)

//#########################################################################
// Consensus Engine Typeclass
//#########################################################################
type EngineI interface {
	// TODO design a suit of `consensus engine interfaces`
}
type Engine struct {
	Name utils.Name
}
