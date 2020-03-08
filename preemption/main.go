// Demonstrates async preemption in 1.14 vs. earlier versions
package main

import (
	"log"
	"runtime"
	"time"
)

func f1() {
	var f float64 = 1.0
	// Going to get a lot of empty loops here in this function
	// They will block in versions earlier than 1.14
	for i := 0; i < 1000000000; i++ {
		f = (f + float64(i)) / f
	}
	log.Println("f1", f)
}

func f2() {
	for {
		time.Sleep(100 * time.Millisecond)
		log.Println("f2")
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// GOMAXPROCS is going to limit the work to just one thread to help illustrate
	// what is going on behind the scenes
	runtime.GOMAXPROCS(1)

	go f2()
	time.Sleep(300 * time.Millisecond)
	go f1()
	select {}
}
