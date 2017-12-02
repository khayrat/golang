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

func fibonacci_() func() int {
    count, fn_1, fn_2 := 0, 0, 1

    return func() int {
        switch {
        case count == 0:
            {
                count++
                return 0
            }
        case count == 1:
            {
                count++
                return 1
            }
        default:
            {
                fn := fn_1 + fn_2
                fn_1 = fn_2
                fn_2 = fn
                return fn
            }
        }
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 15; i++ {
        fmt.Println(f())
    }
}

