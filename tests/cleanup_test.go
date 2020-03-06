package main

import (
	"testing"
)

func TestCleanup(t *testing.T) {
	t.Log("Let's test some stuff")
	t.Cleanup(func() {
		t.Log("I'm cleaning up!")
	})
	t.Fatal("We're calling t.Fatal here. The cleanup function should still run though :)")
}
