package raft

import (
	"net"
	"net/rpc"
	"sync"
)

type Server struct {
	mu *sync.Mutex

	serverId int
	peerIds  []int

	rpcServer *rpc.Server
	listener  net.Listener

	peerClients map[int]*rpc.Client

	wg sync.WaitGroup
}

func NewServer(serverId int, peerIds []int) *Server {
	srv := new(Server)

	srv.serverId = serverId
	srv.peerIds = peerIds
	srv.peerClients = make(map[int]*rpc.Client)

	return srv
}

func (srv *Server) Serve() {
	// todo:
	// 1. initialize rpc
	// 2. connect to all peers
	// 3. start rpc

}
