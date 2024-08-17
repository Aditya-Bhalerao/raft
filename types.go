package main

import (
	"sync"
	"time"
)

type ServerState int

const (
	Follower ServerState = iota
	Candidate
	Leader
	Offline
)

type LogEntry struct {
	Term    int
	Command string
	Index   int
}

type Server struct {
	mu    sync.Mutex
	id    int
	state ServerState

	// Persistent state on all servers:
	currentTerm int
	votedFor    int
	log         []LogEntry

	// Volatile state on all servers:
	commitIndex int
	lastApplied int

	// Volatile state on leaders:
	nextIndex  map[int]int
	matchIndex map[int]int

	electionTimer *time.Timer
	heartbeatChan chan bool
}
