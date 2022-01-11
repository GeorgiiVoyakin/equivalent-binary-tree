package main

import (
	"fmt"
	"sort"
	"time"

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
	var tree1 []int
	var tree2 []int
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// for i := 0; i < 10; i++ {
	// 	x := <-ch1
	// 	tree1 = append(tree1, x)
	// 	x = <-ch2
	// 	tree2 = append(tree2, x)

	// }

	br := false
	time.Sleep(1 * time.Second)
	for {
		select {
		case x := <-ch1:
			tree1 = append(tree1, x)
		case x := <-ch2:
			tree2 = append(tree2, x)
		default:
			br = true
		}

		if br {
			break
		}
	}

	fmt.Println(tree1)
	fmt.Println(tree2)
	sort.Slice(tree1, func(i, j int) bool { return tree1[i] < tree1[j] })
	sort.Slice(tree2, func(i, j int) bool { return tree2[i] < tree2[j] })

	fmt.Println(tree1)
	fmt.Println(tree2)

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
	fmt.Println(Same(tree.New(2), tree.New(2)))
}
