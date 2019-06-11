package poc

import (
	"../consensus"
	"../p2pnet"
	"../utils"

	"fmt"
)

//#########################################################################
// Constraint-Box Typeclass
//#########################################################################

//type constraints []*Constraint
type constraints map[utils.Name]*Constraint
type constraintI interface {
	process()
	forget()
}
type process func()
type forget func()
type Constraint struct {
	name    utils.Name
	Process process
	Forget  forget
}

func makeConstraint(name utils.Name, ci constraintI) *Constraint {
	return &Constraint{name, ci.process, ci.forget}
}

//#########################################################################
// Nodoe Constraint, which respected a specifice `p2pnet` server
//#########################################################################

type Node struct {
	constr     *Constraint
	connectors connectors
	node       p2pnet.Node
}

func MakeNoode(node p2pnet.Node, conn *Connector) Node {
	self := Node{node: node}
	self.constr = makeConstraint(node.Name, &self)
	self.connectors[conn.name] = conn
	conn.Connect(self.constr)
	return self
}

func (self Node) process() {
	// TODO do some hash upgrades

}

func (self Node) forget() {
	for _, conn := range self.connectors {
		conn.Forget()
	}
}

//#########################################################################
// Spreader Constraint, which respected a specifice `blockahain` application
//#########################################################################

type Spreader struct {
	constr     *Constraint
	connectors connectors
	engine     consensus.Engine
}

func MakeSpreader(engine consensus.Engine, conn *Connector) Spreader {
	self := Spreader{engine: engine}
	self.constr = makeConstraint(engine.Name, &self)
	self.connectors[conn.name] = conn
	conn.Connect(self.constr)
	return self
}

func (self Spreader) process() {
	// TODO do some consensus

}

func (self Spreader) forget() {
	for _, conn := range self.connectors {
		conn.Forget()
	}
}

//#########################################################################
// Probe Constraint, which respected a commonly `constraint printer`
//#########################################################################

type Probe struct {
	constr    *Constraint
	connector *Connector
}

func MakeProbe(name utils.Name, conn *Connector) Probe {
	self := Probe{connector: conn}
	self.constr = makeConstraint(name, &self)
	conn.Connect(self.constr)
	return self
}

func (self Probe) process() {
	self.print(self.connector.GetVal())
}

func (self Probe) forget() {
	self.print("?")
}

func (self Probe) print(value interface{}) {
	fmt.Printf("Probe: %s = %v", self.connector.name, value)
}
