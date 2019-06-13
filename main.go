package main

import (
	"fmt"

	"./p2pnet"
	"./poc"
	"./utils"
)

func main() {
	//var a = poc.MakeConnector("a")
	//var b = poc.MakeConnector("b")
	//var c = poc.MakeConnector("c")
	//var d = poc.MakeConnector("d")
	//var e = poc.MakeConnector("e")
	//var f = poc.MakeConnector("f")
	//var g = poc.MakeConnector("g")
	//
	//var node1 = poc.MakeNode(p2pnet.MakeNode("node1", "192.168.1.1"), conn1)
	//var node2 = poc.MakeNode(p2pnet.MakeNode("node2", "192.168.1.2"), conn1)
	//var node3 = poc.MakeNode(p2pnet.MakeNode("node3", "192.168.1.3"), conn1)
	//var node4 = poc.MakeNode(p2pnet.MakeNode("node4", "192.168.1.4"), conn1)
	//
	//var chain1 = poc.MakeBlcokchain(utils.GenesisChain(1, "chain1"), conn1)
	//var chain2 = poc.MakeBlcokchain(utils.GenesisChain(2, "chain2"), conn1)
	//
	//
	//var consensus1 = poc.MakeConsensus(poa.MakePoA("poa"), conn1)
	//var consensus2 = poc.MakeConsensus(pos.MakePoS("pos"), conn1)

	var c1 = poc.MakeConnector("c1")
	var n1 = poc.MakeNode(p2pnet.MakeNode("node1", "192.168.1.1"), c1)

	c1.AddVal(utils.Tx{})

	fmt.Println(c1, n1)



	var v II
	v = S2{}
	fmt.Println(v.foo(), v.bar())

}

type II interface {
	foo() int
	bar() int
}

//empty struct with shared foo() implementation
type S1 struct {
}
func (s S1) foo() int {
	return 123
}

type S2 struct {
	S1
}

func (s S2) bar() int {
	return 456
}