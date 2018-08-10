package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 1000; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(100)
	go say("world")
	say("hello")

}
