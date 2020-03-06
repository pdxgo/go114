package main

import (
	"testing"
	"time"
)

func TestLogStreaming(t *testing.T) {
	t.Logf("Hello! Let's start streaming")
	t.Logf("I'm about to sleep for 1 second")
	time.Sleep(1 * time.Second)
	t.Logf("Done sleeping. Let's sleep for another 2 seconds")
	time.Sleep(2 * time.Second)
	t.Logf("Done again!")
	t.Logf("See how all of these log lines got printed out after each sleep, instead of all at the end? The magic of streaming...")
}
