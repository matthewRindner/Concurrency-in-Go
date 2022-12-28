package main

import (
	"fmt"
	"time"
)

// Channels are a way for go rountines to communicate with each other
// channels also synchronize multiple go routines by waiting for the sender and receiver to fill with data
func main() {
	//declare channel
	c := make(chan string)
	go count("sheep", c)

	//receive message
	//sending and recieving are blocking operations, send will wait for something to send, same with receive
	//iternating over the range of the channel will allow us to keep receiving messages till the channel is closed,
	//so no need for the inside check
	for msg := range c {
		//msg, open := <-c

		//if channel is not open, break out of loop once channel is closed, not checking will result in program never finishing
		// if !open {
		// 	break
		// }
		fmt.Println(msg)
	}
	//when waiting for all the message, program will terminate by deadlock
	//the main() will still wait to receive a message, yet none will ever be sent and the program will never finish
	//go can catch this at runtime, not at compile time cause it doesnt solve the halting problem
	//solve this by closing the chennel

}

// to communicate back to the main go routine, add a channel in the args for count()
// channels have a type so string channels can only pass messages that are strings
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {

		//instead of communicating to the console use the channel
		//<- send message
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	//only close the sender
	//receiver should never close a channel as it doesnt know whether the sender is finshed sening data or not
	close(c)
}
