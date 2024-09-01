// note: need this?

package raft

type Node struct {
	id   int
	addr string
}

var Cluster = []Node{
	{
		id:   1,
		addr: "localhost:10001",
	},
	{
		id:   2,
		addr: "localhost:10002",
	},
	{
		id:   3,
		addr: "localhost:10003",
	},
	{
		id:   4,
		addr: "localhost:10004",
	},
	{
		id:   5,
		addr: "localhost:10005",
	},
}
