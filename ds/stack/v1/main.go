package main

import (
	"fmt"
	stack "lab/goratch/ds/stack/v1/impl"
)

func main() {
	arr := []string{"1", "2", "4", "7", "9"}
	stack1 := stack.NewStack[string]()
	for _, s := range arr {
		fmt.Printf("Push \"%s\" to stack\n", s)
		stack1.Push(s)
	}
	for i := range stack1.Size() {
		val, _ := stack1.Peek()
		fmt.Printf("Peek index %d \"%s\"\n", i, val)
	}
	fmt.Printf("Stack after peek %v \n", stack1)

	for i := range stack1.Size() {
		val, _ := stack1.Pop()
		fmt.Printf("Peek index %d \"%s\"\n", i, val)
	}
	fmt.Printf("Stack after pop %v \n", stack1)
}
