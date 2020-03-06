package main

import (
	"fmt"
)

func gr1() {
	for {
	}
}

func main() {
	go gr1()
	fmt.Println("Hello from main")
}
