/*Develop a microservice for keeping the count of views on a video platform (YouTube). It should have 2 APIs -
API 1: for a given alphanumeric videoId (e.g. 0ReKdcpNyQg) increment the view count
API 2: for a give videoId return number of views

You can store data in memory.

- Run your code with -race flag and call both apis multiple times without facing any error
- Implement unit tests*/

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

	//adding few print statements
	fmt.Println("hello world")
	fmt.Println("hello world2")

	fmt.Println("Views on video2 - ", <-c2)

	vid.Wg.Wait()
	vid2.Wg.Wait()

}
