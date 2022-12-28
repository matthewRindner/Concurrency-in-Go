package main

import (
	"fmt"
	"sync"
	"time"
)

// waits groups are used to solve the need for user input with n amount of go routines

func main() {
	//declare new wait group.
	//wait groups act like a counter for the amount of go routines
	//thus we have 1 go routine to wait for
	//increment of wait group must be done by programmer
	var wg sync.WaitGroup
	wg.Add(1)

	//declearation of anomyous function i avoid passing a pointer to count()
	// as its not the responsibility of count() to decrement the wait group
	go func() {
		count("sheep")
		wg.Done() //decrements wg counter by 1
	}()

	//A bloacking call that will wait tll the wait group counter is 0
	//wait for any unfinshed go routines
	wg.Wait()

}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
