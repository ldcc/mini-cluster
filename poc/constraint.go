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
	snapshoot()
	process(utils.Name, utils.Cv)
	commit(utils.Name, utils.CvSet)
}

//###################################################################################
// Constraint-Box Typeclass
// 一个约束器可能同时被多个调度器所捆绑
// 一个调度器也可能同时关联了多个约束器
//###################################################################################

type Constraint struct {
	Application
	Name      utils.Name
	stores    utils.CvSet
	validated utils.CvSet
	dispatchs dispatchs
}

func makeConstraint(cname utils.Name, app Application, disps ...*Dispatcher) *Constraint {
	dispatchs := make(dispatchs)
	constr := &Constraint{
		Application: app,
		Name:        cname,
		validated:   make(utils.CvSet),
		stores:      make(utils.CvSet),
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

func (constr *Constraint) Propagate(sender utils.Name, cv utils.Cv) {
	if constr.stores.Exist(cv) {
		return
	}

	constr.stores.AddCv(cv)
	constr.process(sender, cv)
	for dname, disp := range constr.dispatchs {
		if dname != sender {
			func() {
				disp.Propagate(constr.Name, cv)
			}()
		}
	}
}

func (constr *Constraint) Commit(sender utils.Name, set utils.CvSet) {
	if len(set) == 0 {
		return
	}

	constr.commit(sender, set)
	//constr.snapshoot()
	for dname, disp := range constr.dispatchs {
		if dname != sender {
			func() {
				disp.Commit(constr.Name, set)
			}()
		}
	}
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

func (probe *Probe) snapshoot() {
	panic("implement me")
}

func (probe *Probe) process(sender utils.Name, cv utils.Cv) {
	probe.print(sender, cv)
}

func (probe *Probe) commit(sender utils.Name, set utils.CvSet) {
	probe.print(sender, set)
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

func (node *Node) snapshoot() {
	panic("implement me")
}

func (node *Node) process(sender utils.Name, cv utils.Cv) {
	if node.valideMessage(sender, cv) {
		fmt.Printf("%s valided %s.%v success!\n", node.constr.Name, sender, cv)
	}
}

func (node *Node) commit(sender utils.Name, set utils.CvSet) {
}

func (node Node) valideMessage(sender utils.Name, cv utils.Cv) bool {
	// TODO do some process
	return true
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

func (chain *Blockchain) snapshoot() {
	panic("implement me")
}

func (chain *Blockchain) process(sender utils.Name, cv utils.Cv) {
	panic("implement me")
}

func (chain *Blockchain) commit(sender utils.Name, set utils.CvSet) {
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

func (cons *Consensus) snapshoot() {
	panic("implement me")
}

func (cons *Consensus) process(sender utils.Name, cv utils.Cv) {
	panic("implement me")
}

func (cons *Consensus) commit(sender utils.Name, set utils.CvSet) {
}
