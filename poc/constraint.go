package poc

import (
	"fmt"
	"mini-cluster/consensus"
	"mini-cluster/p2pnet"
	"mini-cluster/utils"
)

// TODO 在单机环境下一个 Dispatcher 与一个 RPC-Cluster 同余
//      在分布式环境下 dispatchs 需要改成 connection-cluster-pool

type dispatchs map[utils.Name]*Dispatcher

type Application interface {
	propagate(*Dispatcher, utils.Cv)
	process(*Dispatcher, utils.Cv)
	commit(*Dispatcher, utils.Cv)
}

//###################################################################################
// Constraint-Box Typeclass
// 一个约束器可能同时被多个调度器所捆绑
// 一个调度器也可能同时关联了多个约束器
//###################################################################################

type Constraint struct {
	Application
	Name      utils.Name
	validated utils.CvSet
	stores    utils.CvSet
	dispatchs dispatchs
}

func (constr *Constraint) Process(sender *Dispatcher, cv utils.Cv) {
	if constr.stores.Exist(cv) {
		return
	}

	for name, disp := range constr.dispatchs {
		if name != sender.Name {
			constr.propagate(disp, cv)
		}
	}
	constr.process(sender, cv)
}

func (constr *Constraint) Commit(sender *Dispatcher, cv utils.Cv) {
	if constr.stores.Exist(cv) {
		return
	}

	for name, disp := range constr.dispatchs {
		if name != sender.Name {
			disp.Commit(constr.Name, cv)
		}
	}
	constr.commit(sender, cv)
}

func (constr *Constraint) Connect(disp *Dispatcher) {
	if _, ok := constr.dispatchs[disp.Name]; !ok {
		constr.dispatchs[disp.Name] = disp
	}
}

func (constr *Constraint) Disconnect(disp *Dispatcher) {
	if _, ok := constr.dispatchs[disp.Name]; ok {
		delete(constr.dispatchs, disp.Name)
	}
}

func makeConstraint(cname utils.Name, app Application, disps ...*Dispatcher) *Constraint {
	dispatchs := make(dispatchs)
	constr := &Constraint{
		Name:        cname,
		Application: app,
		dispatchs:   dispatchs,
	}
	for _, disp := range disps {
		if disp != nil {
			dispatchs[disp.Name] = disp
			disp.Connect(constr)
		}
	}
	return constr
}

//###################################################################################
// Probe Constraint, which respected a commonly `constraint printer`
//###################################################################################

type Probe struct {
	constr *Constraint
}

func MakeProbe(name utils.Name, disps ...*Dispatcher) *Constraint {
	self := &Probe{}
	self.constr = makeConstraint(name, self, disps...)
	return self.constr
}

func (probe Probe) propagate(*Dispatcher, utils.Cv) {
}

func (probe Probe) process(sender *Dispatcher, cv utils.Cv) {
	probe.print(sender.Name, cv)
}

func (probe Probe) commit(sender *Dispatcher, cv utils.Cv) {
	probe.print(sender.Name, "?")
}

func (probe Probe) print(name utils.Name, msg interface{}) {
	fmt.Printf("Probe: %s \nNew Message: %v\n", name, msg)
}

//###################################################################################
// Node Constraint, which respected a specifice `p2pnet` server
//###################################################################################

type Node struct {
	constr *Constraint
	peer   *p2pnet.Peer
}

func MakeNode(peer *p2pnet.Peer, disps ...*Dispatcher) *Constraint {
	self := &Node{peer: peer}
	self.constr = makeConstraint(peer.Name, self, disps...)
	return self.constr
}

func (node Node) propagate(disp *Dispatcher, msg utils.Cv) {
	disp.SendMessage(msg, node.constr.Name)
}

func (node Node) process(sender *Dispatcher, cv utils.Cv) {
	// TODO do some proccess
}

func (node Node) commit(sender *Dispatcher, cv utils.Cv) {
}

//###################################################################################
// Blcokchain Constraint, which respected a specifice `blockahain` application
//###################################################################################

type Blockchain struct {
	constr *Constraint
	chain  *utils.Chain
}

func MakeBlcokchain(chain *utils.Chain, disps ...*Dispatcher) *Constraint {
	self := &Blockchain{chain: chain}
	self.constr = makeConstraint(utils.Name(chain.RootHash), self, disps...)
	return self.constr
}

func (chain Blockchain) propagate(*Dispatcher, utils.Cv) {
}

func (chain Blockchain) process(sender *Dispatcher, cv utils.Cv) {
	// TODO do some upgrades
}

func (chain Blockchain) commit(sender *Dispatcher, cv utils.Cv) {
}

//###################################################################################
// Consensus Constraint, which respected a specifice `consensus` mechanism
//###################################################################################

type Consensus struct {
	constr *Constraint
	engine *consensus.Engine
}

func MakeConsensus(engine *consensus.Engine, disps ...*Dispatcher) *Constraint {
	self := &Consensus{engine: engine}
	self.constr = makeConstraint(engine.Name, self, disps...)
	return self.constr
}

func (cons Consensus) propagate(*Dispatcher, utils.Cv) {
}

func (cons Consensus) process(sender *Dispatcher, cv utils.Cv) {
	// TODO do some consensus
}

func (cons Consensus) commit(sender *Dispatcher, cv utils.Cv) {
}
