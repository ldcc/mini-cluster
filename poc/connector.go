package poc

import (
	"../utils"
)

//type connectors []*Connector
type connectors map[utils.Name]*Connector
type Connector struct {
	constrs constraints
	stores  utils.TxSet
	value   utils.Tx
	pline   chan *utils.Tx
	name    utils.Name
}

func MakeConnector(name utils.Name) *Connector {
	return &Connector{
		constrs: make(constraints),
		stores:  make(utils.TxSet),
		pline:   make(chan *utils.Tx),
		name:    name,
	}
}

func (self Connector) Connect(constr *Constraint) {
	self.constrs[constr.Name] = constr
	constr.Connect(&self)
}

//func (self Connector) Disconnect(constr *Constraint) {
//	delete(self.constrs, constr.name)
//	constr.Disconnect(&self)
//}

func (self Connector) HasVal() bool {
	return self.stores.HasTx()
}

func (self Connector) GetVal() utils.TxSet {
	return self.stores.Copy()
}

// TODO add mutex lock
func (self Connector) AddVal(tx utils.Tx) {
	self.stores.AddTx(tx)
	//self.pline <- &tx
	self.value = tx
	for _, constr := range self.constrs {
		constr.Process(self.name)
	}
}

//// TODO add mutex lock
func (self Connector) ClsVal() {
	if self.HasVal() {
		self.stores.Clean()
	}
}

// TODO add mutex lock
func (self Connector) Forget(name utils.Name) {
	for cname, constr := range self.constrs {
		if cname != name {
			func() {
				constr.Forget(self.name)
				constr.Process(self.name)
			}()
		}
	}
}
