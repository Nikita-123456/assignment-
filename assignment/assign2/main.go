//------Implement LRU cache-----

package main

import "fmt"

const cap = 2

var mapIt map[int]*node

var head *node
var tail *node

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func create(k int, v int) *node {
	n := node{}
	n.key = k
	n.value = v
	return &n
}

// add node at the front of the linked list
func addNode(newNode *node) {
	temp := head.next
	newNode.next = temp
	newNode.prev = head
	head.next = newNode
	temp.prev = newNode
}

// remove node from the linked list
func deleteNode(delNode *node) {
	delPrev := delNode.prev
	delNext := delNode.next
	delPrev.next = delNext
	delNext.prev = delPrev

}

// brings the corresponding value to the key if present else return -1
// if the key is present then bringing the node at the front of the list and storing new address in the map
func Get(k int) int {
	if _, ok := mapIt[k]; ok {
		resNode := mapIt[k]
		res := resNode.value
		delete(mapIt, k)
		deleteNode(resNode)
		addNode(resNode)
		mapIt[k] = head.next
		return res
	}
	return -1
}

// checks if the key is already present , if yes then delete it from the list and map
// if capacity is full then delete the last node
// at last storing the new node in list and map
func Put(k int, v int) {
	if _, ok := mapIt[k]; ok {
		existingNode := mapIt[k]
		delete(mapIt, k)
		deleteNode(existingNode)
	}
	if len(mapIt) == cap {
		delete(mapIt, tail.prev.key)
		deleteNode(tail.prev)
	}
	addNode(create(k, v))
	mapIt[k] = head.next
}

func main() {

	mapIt = make(map[int]*node)

	//-------------------------------------------------
	//creating 2 nodes head and tail and making them point to each other
	//nil<---head<====>tail--->nil

	head = create(-1, -1)
	tail = create(-1, -1)

	head.next = tail
	head.prev = nil

	tail.prev = head
	tail.next = nil
	//-------------------------------------------------

	Put(1, 10)          // nil, linked list: [1:10]
	Put(2, 20)          // nil, linked list: [2:20,1:10]
	fmt.Println(Get(1)) // 10, linked list: [1:10,2:20]
	Put(3, 30)          // nil, linked list: [3:30,1:10]
	fmt.Println(Get(2)) // -1, linked list: [3:30,1:10]
	Put(4, 40)          // nil, linked list: [4:40,3:30]
	fmt.Println(Get(1)) // -1, linked list: [4:40,3:30]
	fmt.Println(Get(3)) // 30, linked list: [3:30,4:40]

	//-------testing  purpose----------
	/*for i := 0; i < 15; i++ {
		Put(i, i+10)
	}

	for e := head.next; e != tail; e = e.next {
		fmt.Printf("%d\t", e.key)
	}

	Put(5, 15)
	println()
	Put(16, 26)
	println()

	for e := head.next; e != tail; e = e.next {
		fmt.Printf("%d\t", e.key)
	}*/

}
