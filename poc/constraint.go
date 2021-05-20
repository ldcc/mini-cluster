package poc

import (
	"fmt"
	"mini-cluster/consensus"
	"mini-cluster/p2pnet"
	"mini-cluster/utils"
)

type (
	process    func(*Dispatcher)
	commit     func(*Dispatcher)
	connect    func(*Dispatcher)
	disconnect func(*Dispatcher)
	// TODO 在单机环境下一个 Dispatcher 与一个 RPC-Cluster 同余
	//      在分布式环境下 dispatchs 需要改成 connection-cluster-pool
	dispatchs map[utils.Name]*Dispatcher
)

type Application interface {
	propagate(*Dispatcher, *utils.Cv)
	process(*Dispatcher)
	commit(*Dispatcher)
}

//###################################################################################
// Constraint-Box Typeclass
//###################################################################################

type Constraint struct {
	Name       utils.Name
	Process    process
	Commit     commit
	Connect    connect
	Disconnect disconnect
}

// 一个约束器可能同时被多个调度器所捆绑
// 一个调度器也可能同时关联了多个约束器
func makeConstraint(cname utils.Name, app Application, disps ...*Dispatcher) *Constraint {
	dispatchs := make(dispatchs)
	constr := &Constraint{
		Name: cname,
		Process: func(sender *Dispatcher) {
			msg := sender.message
			for name, disp := range dispatchs {
				if name != sender.name {
					app.propagate(disp, &msg)
				}
			}
			app.process(sender)
		},
		Commit: func(sender *Dispatcher) {
			for name, disp := range dispatchs {
				if name != sender.name {
					disp.Commit(cname)
				}
			}
			app.commit(sender)
		},
		Connect: func(disp *Dispatcher) {
			if _, ok := dispatchs[disp.name]; !ok {
				dispatchs[disp.name] = disp
			}
		},
		Disconnect: func(disp *Dispatcher) {
			if _, ok := dispatchs[disp.name]; ok {
				delete(dispatchs, disp.name)
			}
		},
	}
	for _, disp := range disps {
		if disp != nil {
			dispatchs[disp.name] = disp
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

func (probe Probe) propagate(*Dispatcher, *utils.Cv) {
}

func (probe Probe) process(sender *Dispatcher) {
	probe.print(sender.name, sender.GetMessage())
}

func (probe Probe) commit(sender *Dispatcher) {
	probe.print(sender.name, "?")
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

func (node Node) propagate(disp *Dispatcher, msg *utils.Cv) {
	disp.SendMessage(msg, node.constr.Name)
}

func (node Node) process(sender *Dispatcher) {
	// TODO do some proccess

}

func (node Node) commit(sender *Dispatcher) {
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

func (chain Blockchain) propagate(*Dispatcher, *utils.Cv) {
}

func (chain Blockchain) process(sender *Dispatcher) {
	// TODO do some upgrades

}

func (chain Blockchain) commit(sender *Dispatcher) {
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

func (cons Consensus) propagate(*Dispatcher, *utils.Cv) {
}

func (cons Consensus) process(sender *Dispatcher) {
	// TODO do some consensus

}

func (cons Consensus) commit(sender *Dispatcher) {
}
