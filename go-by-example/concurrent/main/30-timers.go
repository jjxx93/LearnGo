package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 2)

	<-time1.C
	fmt.Println("Time 1 expired")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Time 2 expired")
	}()

	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
