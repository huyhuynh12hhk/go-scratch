package main

import (
	"fmt"
	queue "lab/goratch/ds/queue/v1/impl"
)

func main() {
	arr := []string{"1", "2", "4", "7", "9"}
	queue1 := queue.NewQueue[string]()
	for _, s := range arr {
		fmt.Printf("Push \"%s\" to queue\n", s)
		queue1.Enqueue(s)
	}
	for i := range queue1.Size() {
		val, _ := queue1.Peek()
		fmt.Printf("Peek index %d \"%s\"\n", i, val)
	}
	fmt.Printf("Queue after peek %v \n", queue1)

	for i := range queue1.Size() {
		val, _ := queue1.Dequeue()
		fmt.Printf("Peek index %d \"%s\"\n", i, val)
	}
	fmt.Printf("Queue after dequeue %v \n", queue1)
}
