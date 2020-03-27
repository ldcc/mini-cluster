package poc

import (
	"github.com/ldcc/mini-cluster/utils"
)

//type connectors []*Connector
type connectors map[utils.Name]*Connector

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

func (conn *Connector) Connect(constr *Constraint) {
	conn.constrs[constr.Name] = constr
	constr.Connect(conn)
}

//func (self *Connector) Disconnect(constr *Constraint) {
//	delete(self.constrs, constr.name)
//	constr.Disconnect(&self)
//}

func (conn *Connector) HasVal() bool {
	return conn.stores.HasTx()
}

func (conn *Connector) GetVal() utils.TxSet {
	return conn.stores.Copy()
}

// TODO add mutex lock
func (conn *Connector) AddVal(value utils.Cv, adder utils.Name) {
	tx := value.Value.(utils.Tx)
	conn.stores.AddTx(&tx)
	//conn.pline <- &tx
	conn.value = value
	for cname, constr := range conn.constrs {
		if cname != adder {
			constr.Process(conn)
		}
	}
}

//// TODO add mutex lock
func (conn *Connector) ClsVal() {
	if conn.HasVal() {
		conn.stores.Clean()
	}
}

// TODO add mutex lock
func (conn *Connector) Forget(name utils.Name) {
	for cname, constr := range conn.constrs {
		if cname != name {
			func() {
				constr.Forget(conn)
				constr.Process(conn)
			}()
		}
	}
}
