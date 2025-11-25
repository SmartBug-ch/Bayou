package main

import "sync"

// in Bayou we just have writes
// here we have the base for a Write to identify its herkunft und timestamp
type writeID struct {
	CSN            int64
	localTimestamp int64
	serverID       int64
}

// das ist die reprawesentation von einem eigentlichen write mit seiner ID
type write struct {
	ID      writeID
	payload string
}

// we need a logentry type to create the logs itself in the server
type logEntry struct {
	write     write
	committed bool
}

// this would represent a server
type server struct {
	ID         int64
	log        []logEntry
	serverLock sync.Mutex
}
