package tasks

import (
	"log/slog"
	"sync"
	"time"
)

var m = map[string]int{"a": 1}
var lock = sync.RWMutex{}

func ConcurrentMap() {
	slog.Info("")
	slog.Info("======> Concurrent Map")
	// https://stackoverflow.com/questions/36167200/how-safe-are-golang-maps-for-concurrent-read-write-operations

	go Read()
	time.Sleep(1 * time.Second)
	go Write()
	time.Sleep(1 * time.Minute)
}

func Read() {
	for {
		read()
	}
}

func Write() {
	for {
		write()
	}
}

func read() {
	lock.RLock()
	defer lock.RUnlock()
	_ = m["a"]
}

func write() {
	lock.Lock()
	defer lock.Unlock()
	m["b"] = 2
}
