package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	//main go routine
	//will output one and then the other even though c1 is ready to send much sooner than c2, not ideal
	//this is because the code is being blocked by the slower channel, so we are waiting 2 seconds for c2
	/* for {
	* 	fmt.Println(<-c1)
	*	fmt.Println(<-c2)
	*}
	 */

	//the select statment will receive from whichever channel is ready
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
