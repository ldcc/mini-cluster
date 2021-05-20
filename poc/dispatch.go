package poc

import (
	"mini-cluster/utils"
)

// TODO 设计一个接口使得它同时满足 local-dispatch 和 RPC-Cluster
// TODO 实现 RPC 环境的 dispatch: RPC-Cluster

type constraints map[utils.Name]*Constraint

type Dispatcher struct {
	name    utils.Name
	stores  utils.CvSet
	message utils.Cv
	pline   chan *utils.Cv
	constrs constraints
}

func MakeConnector(name utils.Name) *Dispatcher {
	return &Dispatcher{
		name:    name,
		stores:  make(utils.CvSet),
		pline:   make(chan *utils.Cv),
		constrs: make(constraints),
	}
}

func (disp *Dispatcher) Connect(constr *Constraint) {
	if _, ok := disp.constrs[constr.Name]; !ok {
		disp.constrs[constr.Name] = constr
		constr.Connect(disp)
	}
}

func (disp *Dispatcher) Disconnect(constr *Constraint) {
	if _, ok := disp.constrs[constr.Name]; ok {
		delete(disp.constrs, constr.Name)
		constr.Disconnect(disp)
	}
}

func (disp *Dispatcher) IsEmpty() bool {
	return disp.stores.HasCv()
}

func (disp *Dispatcher) GetMessage() utils.Cv {
	return disp.message
}

// TODO add mutex lock
func (disp *Dispatcher) SendMessage(cv *utils.Cv, adder utils.Name) {
	disp.stores.AddCv(cv)
	//disp.pline <- cv
	disp.message = *cv
	for cname, constr := range disp.constrs {
		if cname != adder {
			constr.Process(disp)
		}
	}
}

// TODO add mutex lock
func (disp *Dispatcher) CleanStores() {
	if disp.IsEmpty() {
		disp.stores.Clean()
	}
}

// TODO add mutex lock
func (disp *Dispatcher) Commit(name utils.Name) {
	for cname, constr := range disp.constrs {
		if cname != name {
			func() {
				constr.Commit(disp)
				constr.Process(disp)
			}()
		}
	}
}
