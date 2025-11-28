package main

import (
	"Bayou/LamportClock"
	"sync"
)

// in Bayou we just have writes
// here we have the base for a Write to identify its herkunft und timestamp
// if CSN == -1 then we dont have any commit of this write yet, or global order
type writeStamp struct {
	commitNum      int64
	localTimestamp int64
	serverID       int64
}

// das ist die reprawesentation von einem eigentlichen write mit seiner ID
type write struct {
	ID      writeStamp
	payload string
}

// we need a logentry type to create the logs itself in the server
type logEntry struct {
	write     write
	committed bool
}

// this would represent a server
// we need a lamportclock implementation to use in a server for the timestamp in the writeStamp
// TODO implement a lamport clock to use in each server
type server struct {
	ID          int64
	log         []logEntry
	serverLock  sync.Mutex
	serverClock *LamportClock.Clock
}

func newWrite(serv server, op string) write {
	return write{
		ID: writeStamp{
			commitNum:      -1,
			localTimestamp: 0, // TODO replace with lamport clock of server
			serverID:       serv.ID,
		},
		payload: op,
	}
}

func newServer(id int64) *server {
	return &server{
		ID:          id,
		log:         make([]logEntry, 0),
		serverClock: LamportClock.GetClock(0),
	}
}
