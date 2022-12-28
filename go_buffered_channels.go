package main

import (
	"fmt"
)

func main() {
	//the num refers to th buffered channel, we can put x many thing into a channel
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	//block on send
	msg := <-c
	fmt.Println(msg)
	//the block will deadlock since the send will block till somting is ready to receive
	//but the code nenver prgresses the the reveive line
	//well need to recive in a seperate go routine OR make something called a buffered channel

	msg = <-c
	fmt.Println(msg)

}
