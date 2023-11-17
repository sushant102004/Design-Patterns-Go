package main

import (
	"fmt"
	"sync"
	"time"
)

type DriverPG struct {
	conn string
}

var instance *DriverPG
var lock = &sync.Mutex{}
var once sync.Once

func Connect() *DriverPG {
	// <-- This code is not thread safe.
	if instance == nil {
		instance = &DriverPG{conn: "localhost"}
	}
	return instance
}

// Thread Safe Code
func ConnectThreadSafe() *DriverPG {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = &DriverPG{conn: "localhost"}
	}
	return instance
}

// Clean & Safe
func ConnectOnce() *DriverPG {
	once.Do(func() {
		instance = &DriverPG{conn: "localhost"}
	})
	return instance
}

func main() {
	// The issue with this code is that several goroutines can evaluate the first check and create singleton instance by replacing other instances.
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Println(*Connect(), " - ", i)
		}
	}()

	go func() {
		fmt.Println(*Connect())
	}()

	fmt.Scanln()
}
