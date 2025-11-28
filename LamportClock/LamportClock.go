package LamportClock

import "sync"

type Clock struct {
	id   uint64
	lock sync.RWMutex
	tick uint64
}

type VectorClock struct {
	lock    sync.RWMutex
	vecTick []uint64
	id      uint64
}

// @return return a pointer to a clock struckt, with the given ocunt start
// clock gets inititlized by itself
func GetClock(i uint64, id uint64) *Clock {
	return &Clock{
		tick: i,
		id:   id,
	}
}

// @return gives a pointer to a vectro clock back, with i participants
// @parameter n is the amount of used clocks and index is the place where the own clock value is placed in the vector ( can be dynamicly increased with AddNewVecClock function TODO do this function shoul
// TODO do a sammlung in here, or how should i actualisieren the clocks ?)
func GetVectorClock(n uint64, id uint64) *VectorClock {
	return &VectorClock{
		vecTick: make([]uint64, n),
		id:      id,
	}
}

// set your own clock to the right value
// @return returns the new clockValue

func RecieveAndSet(recieved *Clock, own *Clock) uint64 {
	if recieved.id <= own.id {
		recieved.lock.Lock()
		own.lock.Lock()
		defer recieved.lock.Unlock()
		defer own.lock.Unlock()

	}
	if recieved.tick > own.tick {
		own.tick = recieved.tick + 1
	} else {
		own.tick += 1
	}
	return own.tick
}

func GetClockValueVec(clock *VectorClock) uint64 {
	clock.lock.RLock()
	defer clock.lock.RUnlock()

	return clock.vecTick[clock.id]
}
func GetClockValue(clock *Clock) uint64 {
	clock.lock.RLock()
	defer clock.lock.RUnlock()

	return clock.tick
}

// does the update after the recieving of a meassage, and updates the own value aas wished ( hence makes this vektor to the new meassage vec.)
// returns 1 on succes else 0
func UpdateVectorClock(recieved *VectorClock, own *VectorClock) uint64 {
	// not same lenth update problem with the clocks updat maybe.
	if recieved.id <= own.id {
		recieved.lock.Lock()
		own.lock.Lock()
		defer recieved.lock.Unlock()
		defer own.lock.Unlock()
	} else {
		own.lock.Lock()
		recieved.lock.Lock()
		defer own.lock.Unlock()
		defer recieved.lock.Unlock()
	}

	if len(recieved.vecTick) != len(own.vecTick) {
		return 0
	}
	for i := 0; i < len(recieved.vecTick); i++ {
		own.vecTick[i] = max(recieved.vecTick[i], own.vecTick[i])
	}
	own.vecTick[own.id] += 1
	return 1
}
