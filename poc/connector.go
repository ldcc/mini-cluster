package poc

import (
	"../utils"
)

//type connectors []*Connector
type connectors map[utils.Name]*Connector

//type Connector struct {
//	Name    utils.Name
//	constrs constraints
//}
type Connector struct {
	constrs constraints
	stores  utils.TxSet
	value   utils.Cv
	pline   chan *utils.Cv
	name    utils.Name
}

func MakeConnector(name utils.Name) *Connector {
	return &Connector{
		constrs: make(constraints),
		stores:  make(utils.TxSet),
		pline:   make(chan *utils.Cv),
		name:    name,
	}
}

func (self *Connector) Connect(constr *Constraint) {
	self.constrs[constr.Name] = constr
	constr.Connect(self)
}

//func (self Connector) Disconnect(constr *Constraint) {
//	delete(self.constrs, constr.name)
//	constr.Disconnect(&self)
//}

func (self *Connector) HasVal() bool {
	return self.stores.HasTx()
}

func (self *Connector) GetVal() utils.TxSet {
	return self.stores.Copy()
}

// TODO add mutex lock
func (self *Connector) AddVal(value utils.Cv, adder utils.Name) {
	tx := value.Value.(utils.Tx)
	self.stores.AddTx(&tx)
	//self.pline <- &tx
	self.value = value
	for cname, constr := range self.constrs {
		if cname != adder {
			constr.Process(self)
		}
	}
}

//// TODO add mutex lock
func (self *Connector) ClsVal() {
	if self.HasVal() {
		self.stores.Clean()
	}
}

// TODO add mutex lock
func (self *Connector) Forget(name utils.Name) {
	for cname, constr := range self.constrs {
		if cname != name {
			func() {
				constr.Forget(self)
				constr.Process(self)
			}()
		}
	}
}
