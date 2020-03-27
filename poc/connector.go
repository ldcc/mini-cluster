package poc

import (
	"github.com/ldcc/mini-cluster/utils"
)

type constraints map[utils.Name]*Constraint

type Dispatch struct {
	name    utils.Name
	stores  utils.CvSet
	message utils.Cv
	pline   chan *utils.Cv
	constrs constraints
}

func MakeConnector(name utils.Name) *Dispatch {
	return &Dispatch{
		name:    name,
		stores:  make(utils.CvSet),
		pline:   make(chan *utils.Cv),
		constrs: make(constraints),
	}
}

func (dispatch *Dispatch) Connect(constr *Constraint) {
	if _, ok := dispatch.constrs[constr.Name]; !ok {
		dispatch.constrs[constr.Name] = constr
		constr.Connect(dispatch)
	}
}

func (dispatch *Dispatch) Disconnect(constr *Constraint) {
	if _, ok := dispatch.constrs[constr.Name]; !ok {
		delete(dispatch.constrs, constr.Name)
		constr.Disconnect(dispatch)
	}
}

func (dispatch *Dispatch) IsEmpty() bool {
	return dispatch.stores.HasCv()
}

func (dispatch *Dispatch) GetMessage() utils.Cv {
	return dispatch.message
}

// TODO add mutex lock
func (dispatch *Dispatch) SendMessage(cv *utils.Cv, adder utils.Name) {
	dispatch.stores.AddCv(cv)
	//dispatch.pline <- cv
	dispatch.message = *cv
	for cname, constr := range dispatch.constrs {
		if cname != adder {
			constr.Process(dispatch)
		}
	}
}

// TODO add mutex lock
func (dispatch *Dispatch) CleanStores() {
	if dispatch.IsEmpty() {
		dispatch.stores.Clean()
	}
}

// TODO add mutex lock
func (dispatch *Dispatch) Commit(name utils.Name) {
	for cname, constr := range dispatch.constrs {
		if cname != name {
			func() {
				constr.Commit(dispatch)
				constr.Process(dispatch)
			}()
		}
	}
}
