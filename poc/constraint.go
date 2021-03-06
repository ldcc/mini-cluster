package poc

import (
	"github.com/ldcc/mini-cluster/consensus"
	"github.com/ldcc/mini-cluster/p2pnet"
	"github.com/ldcc/mini-cluster/utils"

	"fmt"
)

type (
	process    func(*Dispatch)
	commit     func(*Dispatch)
	connect    func(*Dispatch)
	disconnect func(*Dispatch)
	// TODO 在单机环境下一个 Dispatch 与一个 RPC-Cluster 同余
	//      在分布式环境下 dispatchs 需要改成 connection-cluster-pool
	dispatchs map[utils.Name]*Dispatch
)

//###################################################################################
// Constraint-Box Typeclass
//###################################################################################

type Application interface {
	propogate(*Dispatch, *utils.Cv)
	process(*Dispatch)
	commit(*Dispatch)
}

type Constraint struct {
	Name       utils.Name
	Process    process
	Commit     commit
	Connect    connect
	Disconnect disconnect
}

func makeConstraint(cname utils.Name, app Application, disps ...*Dispatch) *Constraint {
	dispatchs := make(dispatchs)
	constr := &Constraint{
		Name: cname,
		Process: func(sender *Dispatch) {
			msg := sender.message
			for name, disp := range dispatchs {
				if name != sender.name {
					app.propogate(disp, &msg)
				}
			}
			app.process(sender)
		},
		Commit: func(sender *Dispatch) {
			for name, disp := range dispatchs {
				if name != sender.name {
					disp.Commit(cname)
				}
			}
			app.commit(sender)
		},
		Connect: func(disp *Dispatch) {
			if _, ok := dispatchs[disp.name]; !ok {
				dispatchs[disp.name] = disp
			}
		},
		Disconnect: func(disp *Dispatch) {
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

func MakeProbe(name utils.Name, disps ...*Dispatch) *Constraint {
	self := &Probe{}
	self.constr = makeConstraint(name, self, disps...)
	return self.constr
}

func (probe Probe) propogate(*Dispatch, *utils.Cv) {
}

func (probe Probe) process(sender *Dispatch) {
	probe.print(sender.name, sender.GetMessage())
}

func (probe Probe) commit(sender *Dispatch) {
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

func MakeNode(peer *p2pnet.Peer, disps ...*Dispatch) *Constraint {
	self := &Node{peer: peer}
	self.constr = makeConstraint(peer.Name, self, disps...)
	return self.constr
}

func (node Node) propogate(disp *Dispatch, msg *utils.Cv) {
	disp.SendMessage(msg, node.constr.Name)
}

func (node Node) process(sender *Dispatch) {
	// TODO do some proccess

}

func (node Node) commit(sender *Dispatch) {
}

//###################################################################################
// Blcokchain Constraint, which respected a specifice `blockahain` application
//###################################################################################

type Blockchain struct {
	constr *Constraint
	chain  *utils.Chain
}

func MakeBlcokchain(chain *utils.Chain, disps ...*Dispatch) *Constraint {
	self := &Blockchain{chain: chain}
	self.constr = makeConstraint(utils.Name(chain.RootHash), self, disps...)
	return self.constr
}

func (chain Blockchain) propogate(*Dispatch, *utils.Cv) {
}

func (chain Blockchain) process(sender *Dispatch) {
	// TODO do some upgrades

}

func (chain Blockchain) commit(sender *Dispatch) {
}

//###################################################################################
// Consensus Constraint, which respected a specifice `consensus` mechanism
//###################################################################################

type Consensus struct {
	constr *Constraint
	engine *consensus.Engine
}

func MakeConsensus(engine *consensus.Engine, disps ...*Dispatch) *Constraint {
	self := &Consensus{engine: engine}
	self.constr = makeConstraint(engine.Name, self, disps...)
	return self.constr
}

func (cons Consensus) propogate(*Dispatch, *utils.Cv) {
}

func (cons Consensus) process(sender *Dispatch) {
	// TODO do some consensus

}

func (cons Consensus) commit(sender *Dispatch) {
}
