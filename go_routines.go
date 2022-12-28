package main

import (
	"fmt"
	"time"
)

func main() {
	//Currently, fish will never execute as sheep never finishes
	// count("sheep")
	// count("fish")

	//program wont wait for sheep to finish
	//go run this program in the backround, then immedialty run the next line
	//main is technically a go routine
	go count("sheep")
	go count("fish")

	//if both functions are go routines, the program will perform the next line
	//since there is no other operation, the program exits

	//the program terminates after 2 seconds for the same reasons

	//time.Sleep(time.Second * 2)

	//this will top the program given user input, not useful to rely on user input
	//this is a blocking call

	//fmt.Scanln()
}

func count(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500) //sleep for half a second
	}
}
