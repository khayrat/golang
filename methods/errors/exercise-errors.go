package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))    
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)    
    }
    
    z := 1.0
    epsilon := 0.000000001

    for  {
        if (math.Abs(x - z*z)) < epsilon {
            return z, nil
        }
        z -= (z*z - x) / (2 * z)
        fmt.Println(z)
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}

