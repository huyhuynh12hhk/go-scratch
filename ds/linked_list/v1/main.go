package main

import (
	"fmt"
	linked_list "lab/goratch/ds/linked_list/v1/impl"
)

func main() {
	var s = []string{"foo", "bar", "zoo"}

	lst := linked_list.FromArray(s)
	fmt.Println("list:", lst.Values())
}
