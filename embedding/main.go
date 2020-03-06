package main

type interface1 interface {
	Hello()
}

type interface2 interface {
	Hello()
}

type interface3 interface {
	interface1
	interface2
}

type interface4 interface {
	Hello()
	interface1
	interface2
}

func main() {
}
