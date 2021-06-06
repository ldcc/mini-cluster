package poc

import (
	"mini-cluster/utils"
)

// TODO 设计一个接口使得它同时满足 local-dispatch 和 RPC-Cluster
// TODO 实现 RPC 环境的 dispatch: RPC-Cluster

type constraints map[utils.Name]*Constraint

type Dispatcher struct {
	Name      utils.Name
	stores    utils.CvSet
	validated utils.CvSet
	constrs   constraints
}

func MakeDispatcher(name utils.Name) *Dispatcher {
	return &Dispatcher{
		Name:      name,
		stores:    make(utils.CvSet),
		validated: make(utils.CvSet),
		constrs:   make(constraints),
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

func (disp *Dispatcher) Empty() bool {
	return disp.stores.Empty()
}

func (disp *Dispatcher) Propagate(sender utils.Name, cv utils.Cv) {
	if disp.stores.Exist(cv) {
		return
	}

	disp.stores.AddCv(cv)
	for cname, constr := range disp.constrs {
		if cname != sender {
			func() {
				constr.Propagate(disp.Name, cv)
			}()
		}
	}
}

// TODO add mutex lock
func (disp *Dispatcher) CleanStores() {
	if !disp.Empty() {
		disp.stores.Clean()
	}
}

// TODO add mutex lock
func (disp *Dispatcher) Commit(sender utils.Name, set utils.CvSet) {
	for cname, constr := range disp.constrs {
		if cname != sender {
			func() {
				constr.Commit(disp.Name, set)
			}()
		}
	}
}
