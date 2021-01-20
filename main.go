package main

import (
	"fmt"
	"github.com/rprajapati0067/raviprajapati-assignment-golang-se/assignment_1"
	"github.com/rprajapati0067/raviprajapati-assignment-golang-se/assignment_2"
)

func main() {

	bst := &assignment_1.BinarySearchTree{}
	eg := []int{2, 7, 1, 5}
	for _, val := range eg {
		bst.Insert(val)
	}
	fmt.Println("**********Assignment 1**********")
	fmt.Println("Printing InOrder:", bst.InOrder())
	fmt.Println("Printing PreOrder:", bst.PreOrder())
	fmt.Println("Printing PostOrder:", bst.PostOrder())

	fmt.Println("**********Assignment 2**********")
	fmt.Printf("Output: %d",
		assignment_2.HouseRob([]int{1, 2, 3, 1}))
}
