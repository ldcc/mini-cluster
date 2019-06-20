package poc

import (
	"../consensus"
	"../p2pnet"
	"../utils"
	"fmt"
)

//###################################################################################
// Constraint-Box Typeclass
//###################################################################################

//type constraints []*Constraint
type constraints map[utils.Name]*Constraint
type constraintI interface {
	process(utils.Name)
	forget(utils.Name)
}
type process func(utils.Name)
type forget func(utils.Name)
type connect func(*Connector) // Implicit `Function Side Effect`
type Constraint struct {
	Name       utils.Name
	connectors connectors
	Process    process
	Forget     forget
	Connect    connect
}

func makeConstraint(name utils.Name, ci constraintI, conns ...*Connector) *Constraint {
	connectors := make(connectors)
	self := &Constraint{name, connectors, ci.process,
		func(name utils.Name) {
			for cname, conn := range connectors {
				if cname != name {
					conn.Forget(name)
				}
			}
			ci.forget(name)
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

func (self Probe) process(name utils.Name) {
	self.print(self.constr.connectors[name])
}

func (self Probe) forget(name utils.Name) {
	self.print("?")
}

func (self Probe) print(value interface{}) {
	conn := value.(*Connector)
	fmt.Printf("Probe: %s stores: %v", conn.name, conn.GetVal())
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

func (self Node) process(name utils.Name) {
	sender := self.constr.connectors[name]
	for cname, conn := range self.constr.connectors {
		if cname != name {
			conn.AddVal(sender.value)
		}
	}
	// TODO do some proccess

}

func (self Node) forget(name utils.Name) {

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

func (self Blockchain) process(name utils.Name) {
	// TODO do some upgrades

}

func (self Blockchain) forget(name utils.Name) {

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

func (self Consensus) process(name utils.Name) {
	// TODO do some consensus

}

func (self Consensus) forget(name utils.Name) {

}
