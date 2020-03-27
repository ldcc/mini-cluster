package p2pnet

import (
	"net"

	utils "github.com/ldcc/mini-cluster/utils"
)

type Peer struct {
	Name utils.Name
	Addr *net.TCPAddr
}

func MakePeer(id, addr string) *Peer {
	var tcpaddr, _ = net.ResolveTCPAddr("tcp", addr)
	return &Peer{Name: utils.Name(id), Addr: tcpaddr}
}
