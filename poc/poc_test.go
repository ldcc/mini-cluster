package poc

import (
	"mini-cluster/p2pnet"
	"mini-cluster/utils"
	"testing"
)

func TestPoC(t *testing.T) {
	var tx = utils.Cv{Hash: "tx1", Value: utils.Tx{}}

	var disp1 = MakeDispatcher("disp1")
	var disp2 = MakeDispatcher("disp2")
	var disp3 = MakeDispatcher("disp3")

	var n = MakeNode(p2pnet.MakePeer("node1", ":9001"), disp1, disp2)
	var _ = MakeNode(p2pnet.MakePeer("node2", ":9002"), disp1)

	var _ = MakeNode(p2pnet.MakePeer("node3", ":9003"), disp1, disp2)
	var _ = MakeNode(p2pnet.MakePeer("node4", ":9004"), disp2)

	//var _ = MakeProbe("p1", disp1, disp2)

	n.Connect(disp3)

	//n.Propagate("", tx)
	disp1.Propagate("", tx)
}
