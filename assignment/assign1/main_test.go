package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	id := "check"

	vid := NewVideo()

	vid.Name[id] = 0

	for i := 0; i < 10; i++ {
		vid.Wg.Add(1)

		go vid.IncViews(id)
	}

	vid.Wg.Wait()

	c := make(chan int)

	vid.Wg.Add(1)

	go vid.GetViews(id, c)

	views := <-c
	if views != 10 {
		fmt.Println("Was expecting 10 views")
	} else {
		fmt.Println("Right !! ")
	}

	vid.Wg.Wait()

}
