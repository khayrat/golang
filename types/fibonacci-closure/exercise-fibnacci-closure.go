package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    x, y := 0, 1
    
    return func() int {
        n := x
        x, y = y, x+y
        
        return n
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 15; i++ {
        fmt.Println(f())
    }
}

