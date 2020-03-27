package poc

import (
	"github.com/ldcc/mini-cluster/consensus"
	"github.com/ldcc/mini-cluster/p2pnet"
	"github.com/ldcc/mini-cluster/utils"

	"fmt"
)

//###################################################################################
// Constraint-Box Typeclass
//###################################################################################

//type constraints []*Constraint
type constraints map[utils.Name]*Constraint
type constraintI interface {
	propogate(*Connector, *utils.Cv)
	process(*Connector)
	forget(*Connector)
}
type process func(*Connector)
type forget func(*Connector)
type connect func(*Connector) // Implicit `Function Side Effect`
type Constraint struct {
	Name utils.Name
	Process process
	Forget  forget
	Connect connect
}

func makeConstraint(name utils.Name, ci constraintI, conns ...*Connector) *Constraint {
	connectors := make(connectors)
	self := &Constraint{
		name,
		func(sender *Connector) {
			value := sender.value // Copy new value
			for cname, conn := range connectors {
				if cname != sender.name {
					ci.propogate(conn, &value)
				}
			}
			ci.process(sender)
		},
		func(sender *Connector) {
			for cname, conn := range connectors {
				if cname != sender.name {
					conn.Forget(name)
				}
			}
			ci.forget(sender)
		},
		func(conn *Connector) {
			connectors[conn.name] = conn
		},
	}
	for _, conn := range conns {
		if conn != nil {
			connectors[conn.name] = conn
			conn.Connect(self)
		}
	}
	return self
}

//###################################################################################
// Probe Constraint, which respected a commonly `constraint printer`
//###################################################################################

type Probe struct {
	constr *Constraint
}

func MakeProbe(name utils.Name, conns ...*Connector) *Constraint {
	self := &Probe{}
	self.constr = makeConstraint(name, self, conns...)
	return self.constr
}

func (probe Probe) propogate(*Connector, *utils.Cv) {
}

func (probe Probe) process(sender *Connector) {
	probe.print(sender.name, sender.GetVal())
}

func (probe Probe) forget(sender *Connector) {
	probe.print(sender.name, "?")
}

func (probe Probe) print(name utils.Name, value interface{}) {
	fmt.Printf("Probe: %s stores: %v\n", name, value)
}

//###################################################################################
// Node Constraint, which respected a specifice `p2pnet` server
//###################################################################################

type Node struct {
	constr *Constraint
	node   *p2pnet.Node
}

func MakeNode(node *p2pnet.Node, conns ...*Connector) *Constraint {
	self := &Node{node: node}
	self.constr = makeConstraint(node.Name, self, conns...)
	return self.constr
}

func (node Node) propogate(conn *Connector, value *utils.Cv) {
	//tx := value.(*utils.Tx)
	conn.AddVal(*value, node.constr.Name)
}

func (node Node) process(sender *Connector) {
	// TODO do some proccess

}

func (node Node) forget(sender *Connector) {
}

//###################################################################################
// Blcokchain Constraint, which respected a specifice `blockahain` application
//###################################################################################

type Blockchain struct {
	constr *Constraint
	chain  *utils.Chain
}

func MakeBlcokchain(chain *utils.Chain, conns ...*Connector) *Constraint {
	self := &Blockchain{chain: chain}
	self.constr = makeConstraint(utils.Name(chain.RootHash), self, conns...)
	return self.constr
}

func (chain Blockchain) propogate(*Connector, *utils.Cv) {
}

func (chain Blockchain) process(sender *Connector) {
	// TODO do some upgrades

}

func (chain Blockchain) forget(sender *Connector) {
}

//###################################################################################
// Consensus Constraint, which respected a specifice `consensus` mechanism
//###################################################################################

type Consensus struct {
	constr *Constraint
	engine *consensus.Engine
}

func MakeConsensus(engine *consensus.Engine, conns ...*Connector) *Constraint {
	self := &Consensus{engine: engine}
	self.constr = makeConstraint(engine.Name, self, conns...)
	return self.constr
}

func (cons Consensus) propogate(*Connector, *utils.Cv) {
}

func (cons Consensus) process(sender *Connector) {
	// TODO do some consensus

}

func (cons Consensus) forget(sender *Connector) {
}
