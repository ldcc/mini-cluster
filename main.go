package main

import (
	"github.com/ldcc/mini-cluster/p2pnet"
	"github.com/ldcc/mini-cluster/poc"
	"github.com/ldcc/mini-cluster/utils"
)

func main() {
	var cv = &utils.Cv{Hash: "tx1", Value: utils.Tx{}}
	var conn1 = poc.MakeConnector("conn1")
	var conn2 = poc.MakeConnector("conn1")
	var n = poc.MakeNode(p2pnet.MakePeer("node1", ":9001"), conn1, conn2)
	var _ = poc.MakeNode(p2pnet.MakePeer("node2", ":9002"), conn1)

	var _ = poc.MakeNode(p2pnet.MakePeer("node3", ":9003"), conn2)
	var _ = poc.MakeNode(p2pnet.MakePeer("node4", ":9004"), conn2)

	poc.MakeProbe("p1", conn1, conn2)

	conn1.SendMessage(cv, n.Name)

}
