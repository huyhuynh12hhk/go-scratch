package main

import (
	"fmt"
	btree "lab/goratch/ds/btree/v1/impl"
)

func main() {
	// str := "Let transform this sentence to a tree then search it"
	str := "G C D A B H I K E F Z"
	// 					D
	// 		B						H
	// A		C			F				K
	// 					E				I		Z
	// nums := []int{2,4,5,3,6,7,6}

	bt1 := btree.FromAString(str)
	// bt1 := btree.FromValueArray(nums)

	in := bt1.ToInOrderTraversalKeys()
	pre := bt1.ToPreOrderTraversalKeys()
	post := bt1.ToPostOrderTraversalKeys()
	fmt.Printf("In Order Value: %+v\n", in)
	fmt.Printf("Pre Order Value: %+v\n", pre)
	fmt.Printf("Post Order Value: %+v\n", post)

	fmt.Println("\nIn Order Nodes:")
	inFull := bt1.ToInOrderTraversalNodeArray()
	for _, n := range inFull {
		fmt.Printf("key: %v | value: %+v | height: %d\n", n.Key, n.Value, n.Height())
	}

	term := "B"
	// term = "it"
	rs, ok := bt1.Search(term)

	if !ok {
		panic(fmt.Sprintf("Error when try to search \"%s\" val %+v", term, rs))
	}

	fmt.Printf("Found: %s (%s) at height %d\n", rs.Value, rs.Key, rs.Height())

}
