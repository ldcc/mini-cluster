package main

import (
	"./consensus/poa"
	"./consensus/pos"
	"./p2pnet"
	"./poc"
	"./utils"
)

func main() {
	var a = poc.MakeConnector("a")
	var b = poc.MakeConnector("b")
	var c = poc.MakeConnector("c")
	var d = poc.MakeConnector("d")
	var e = poc.MakeConnector("e")
	var f = poc.MakeConnector("f")
	var g = poc.MakeConnector("g")

	var n1 = poc.MakeNode(p2pnet.MakeNode("node1", "192.168.1.1"), a)
	var n2 = poc.MakeNode(p2pnet.MakeNode("node2", "192.168.1.2"), b)
	//var n3 = poc.MakeNode(p2pnet.MakeNode("node3", "192.168.1.3"), conn1)
	var n4 = poc.MakeNode(p2pnet.MakeNode("node4", "192.168.1.4"), e)
	var n5 = poc.MakeNode(p2pnet.MakeNode("node5", "192.168.1.5"), g)
	g.Connect(n4)

	poc.MakeBlcokchain(utils.GenesisChain(1, "chain1"), conn1)
	poc.MakeBlcokchain(utils.GenesisChain(2, "chain2"), conn1)
	//poc.MakeBlcokchain(utils.GenesisChain(3, "chain3"), conn1)

	poc.MakeConsensus(poa.MakePoA("poa"), conn1)
	poc.MakeConsensus(pos.MakePoS("pos"), conn1)
	//poc.MakeConsensus(pow.MakePoW("pow"), conn1)

	//var tx1 = utils.Tx{Hash: "tx1"}
	//var c1 = poc.MakeConnector("c1")
	//poc.MakeProbe("p1", c1)
	//poc.MakeNode(p2pnet.MakeNode("node1", "192.168.1.148"), c1)
	//c1.AddVal(tx1)

	//var v II
	//v = S2{}
	//fmt.Println(v.foo(), v.bar())

}

//
//type II interface {
//	foo() int
//	bar() int
//}
//
////empty struct with shared foo() implementation
//type S1 struct {
//}
//
//func (s S1) foo() int {
//	return 123
//}
//
//type S2 struct {
//	S1
//}
//
//func (s S2) bar() int {
//	return 456
//}
