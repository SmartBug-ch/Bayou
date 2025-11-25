package main

// in Bayou we just have writes here ist the ground structure of a write
type write struct {
	CSN            int64
	localTimestamp int64
	serverID       int64
}

// we need a logentry type to create the logs itself in the server
type logEntry struct {
	write     write
	committed bool
}

// this would represent a server
type server struct {
	ID  int64
	log []logEntry
}

func main() {

}
