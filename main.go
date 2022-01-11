package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	tree1 := make([]int, 5)
	tree2 := make([]int, 5)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		select {
		case x := <-ch1:
			tree1 = append(tree1, x)
		case x := <-ch2:
			tree2 = append(tree2, x)
		default:
			break
		}
	}

	sort.Slice(tree1, func(i, j int) bool { return i < j })
	sort.Slice(tree2, func(i, j int) bool { return i < j })
	if len(tree1) != len(tree2) {
		return false
	}

	for i := 0; i < len(tree1); i++ {
		if tree1[i] != tree2[i] {
			return false
		}
	}
	return true
}

func main() {
	// Same(tree.New(1), tree.New(1))
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		default:
			break
		}
	}
}
