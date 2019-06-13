package poc

import (
	"../utils"
)

//type connectors []*Connector
type connectors map[utils.Name]*Connector
type Connector struct {
	constrs constraints
	//value   utils.Tx
	value utils.TxSet
	name  utils.Name
}

func MakeConnector(name utils.Name) *Connector {
	return &Connector{name: name, value: make(utils.TxSet)}
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
	return self.value.HasTx()
}

func (self Connector) GetVal() utils.TxSet {
	return self.value.Copy()
}

// TODO add mutex lock
func (self Connector) AddVal(tx utils.Tx) {
	self.value.AddTx(tx)
	for _, constr := range self.constrs {
		go constr.Process(self.name)
	}
}

// TODO add mutex lock
func (self Connector) ClsVal() {
	if self.HasVal() {
		self.value.Clean()
	}
}

// TODO add mutex lock
func (self Connector) Forget(name utils.Name) {
	for cname, constr := range self.constrs {
		if cname != name {
			go func() {
				constr.Forget(self.name)
				constr.Process(self.name)
			}()
		}
	}
}
