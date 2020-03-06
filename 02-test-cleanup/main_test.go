package main

import (
	"testing"
)

func TestCleanup(t *testing.T) {
	t.Log("Let's test some stuff")
	t.Cleanup(func() {
		t.Log("I'm cleaning up!")
	})
}
