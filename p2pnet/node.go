package p2pnet

import (
	"net"

	"github.com/ldcc/mini-cluster/utils"
)

type Node struct {
	Name utils.Name
	Addr *net.TCPAddr
}

func MakeNode(id, addr string) *Node {
	var tcpaddr, _ = net.ResolveTCPAddr("tcp", addr)
	return &Node{utils.Name(id), tcpaddr}
}
