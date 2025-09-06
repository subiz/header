package header

import (
	"hash/fnv"
	"sync"
	"time"
)

const shardCount = 128

// refCountedMutex is a mutex with a reference counter.
type refCountedMutex struct {
	sync.Mutex
	refCount int
}

var _locks [shardCount]*sync.Mutex
var _shards [shardCount]map[string]*refCountedMutex

// init initializes the shards.
func init() {
	for i := 0; i < shardCount; i++ {
		_shards[i] = map[string]*refCountedMutex{}
		_locks[i] = &sync.Mutex{}
	}

	go func() {
		for {
			time.Sleep(30 * time.Minute)
			// shrink map memory
			for i := 0; i < shardCount; i++ {
				_locks[i].Lock()
				newm := map[string]*refCountedMutex{}
				for k, v := range _shards[i] {
					if v.refCount > 0 {
						newm[k] = v
					}
				}
				_shards[i] = newm
				_locks[i].Unlock()
			}
		}
	}()
}

// getShard returns the correct shard for a given key using FNV-1a hashing.
func getShard(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32() % shardCount)
}

// KLock locks a given key. It returns an unlock function.
func KLock(key string) func() {
	shard := getShard(key)
	lock := _locks[shard]
	lock.Lock()
	mapshard := _shards[shard]
	m, has := mapshard[key]
	if !has {
		m = &refCountedMutex{}
		mapshard[key] = m
	}
	m.refCount++
	lock.Unlock()
	m.Lock()

	return func() {
		lock.Lock()
		mapshard := _shards[shard]
		m.refCount--
		if m.refCount == 0 {
			delete(mapshard, key)
		}
		m.Unlock()
		lock.Unlock()
	}
}
