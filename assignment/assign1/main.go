package main

import (
	"fmt"
	"sync"
)

type Videos struct {
	Lock sync.Mutex
	Wg   sync.WaitGroup
	Name map[string]int
}

func NewVideo() Videos {
	return Videos{
		Name: make(map[string]int),
	}
}

// increase the views
func (v *Videos) IncViews(Id string) {
	defer v.Wg.Done()

	v.Lock.Lock()
	v.Name[Id]++
	v.Lock.Unlock()
}

// get the views
func (v *Videos) GetViews(Id string, c chan int) {
	defer v.Wg.Done()
	v.Lock.Lock()
	c <- v.Name[Id]
	v.Lock.Unlock()
}

func main() {
	id := "abcdef"
	id2 := "xyz"

	vid := NewVideo()
	vid2 := NewVideo()

	vid.Name[id] = 0
	vid2.Name[id2] = 0

	//calling go routines for two videos 100 times
	for i := 0; i < 100; i++ {
		vid.Wg.Add(1)
		vid2.Wg.Add(1)
		go vid.IncViews(id)
		go vid2.IncViews(id2)
	}

	vid.Wg.Wait()
	vid2.Wg.Wait()

	c := make(chan int)
	c2 := make(chan int)

	vid.Wg.Add(1)
	vid2.Wg.Add(1)

	go vid.GetViews(id, c)
	go vid2.GetViews(id2, c2)

	fmt.Println("Views on video1 - ", <-c)
	fmt.Println("Views on video2 - ", <-c2)

	vid.Wg.Wait()
	vid2.Wg.Wait()

}
