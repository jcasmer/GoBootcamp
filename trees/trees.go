// Exercise: Equivalent Binary Trees
// 1. Implement the Walk function.

// 2. Test the Walk function.

// The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

// Create a new channel ch and kick off the walker:

// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

// 4. Test the Same function.

// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

// The documentation for Tree can be found here.

package main

import ( 
	"golang.org/x/tour/tree"
	"fmt"
)
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	
	//fmt.Print(t.Left)	
	//fmt.Print(t.Value)
	//fmt.Print(t.Right)	
	walkTree(t, ch)
	close(ch)
	//ch <- t.Value
}

func walkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
	
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(tree.New(1), ch1)
	go Walk(tree.New(1), ch2)
	for i :=  range ch1 {
		if i != <- ch2 {
			return false
		}		
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Print(Same(tree.New(1), tree.New(1)))
	fmt.Print(Same(tree.New(1), tree.New(2)))
	
}