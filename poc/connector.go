package poc

import (
	"../utils"
)

//type connectors []*Connector
type connectors map[utils.Name]*Connector
type Connector struct {
	constrs constraints
	name    utils.Name
	chain   utils.Chain
	value   *utils.TxList
}

func MakeConnector(name utils.Name) Connector {
	return Connector{name: name, value: new(utils.TxList)}
}

func (self Connector) HasVal() bool {
	return self.value.HasTx()
}

func (self Connector) GetVal() utils.TxList {
	if self.HasVal() {
		return *self.value
	} else {
		return make(utils.TxList, 0)
	}
}

// TODO add mutex lock
func (self Connector) AddVal(tx utils.Tx) {
	self.value.AddTx(tx)
	for _, constr := range self.constrs {
		go constr.Process()
	}
}

// TODO add mutex lock
func (self Connector) ClsVal() {
	if self.HasVal() {
		self.value.Clean()
	}
}

func (self Connector) Connect(constr *Constraint) {
	self.constrs[constr.name] = constr
	//self.constrs = append(self.constrs, constr)
}

func (self Connector) Forget() {
	for name, constr := range self.constrs {
		if name == constr.name {
			constr.Forget()
		}
	}
}

