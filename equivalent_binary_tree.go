//File_name: equivalent_binary_tree.go
//Author: Wenbin Xiao
//Description: http://tour.golang.org/#72

package main

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	c1, ok1 := <-ch1
	c2, ok2 := <-ch2
	for ok1 && ok2 && c1 == c2 {
		c1, ok1 = <-ch1
		c2, ok2 = <-ch2

	}
	return ok1 == ok2 && c1 == c2
}

func main() {
	fmt.Println(Same(tree.New(100), tree.New(100)))
}
