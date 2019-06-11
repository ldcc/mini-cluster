package main

import (
	"fmt"

	"./p2pnet"
	"./poc"
)

func main() {
	var conn = poc.MakeConnector("conn1")

	var node1 = p2pnet.MakeNode("node1", "192.168.1.148")
	var spread1 = poc.MakeNoode(node1, &conn)

	fmt.Println(spread1)
}
