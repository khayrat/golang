package main

import "golang.org/x/tour/tree"

import (
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    var walker func (*tree.Tree)
    walker = func(t *tree.Tree) {
        if t == nil {
            return
        }

        walker(t.Left)
        ch <- t.Value
        walker(t.Right)
    }
    
    walker(t)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1 := make(chan int)
    c2 := make(chan int)

    go Walk(t1, c1)
    go Walk(t2, c2)

    for {
        v1,ok1 := <-c1 
        v2,ok2 := <-c2
        
        fmt.Printf("v1: %v == v2: %v, %v\n", v1, v2, v1 == v2)

        if v1 != v2 || ok1 != ok2 {
            return false
        }
        if !ok1 {
            break    
        }
    }
    return true
}

func main() {
    t1 := tree.New(1)
    t2 := tree.New(1)

    fmt.Printf("t1: %v\nt2: %v\n", t1, t2)
    fmt.Println(Same(t1, t2))
}

