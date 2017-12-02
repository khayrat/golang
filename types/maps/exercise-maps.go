package main

import (
    "golang.org/x/tour/wc"
    //"fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    wc := make(map[string]int)
    
    for _,v := range(strings.Fields(s)) {
        wc[v] = wc[v] + 1
    }
    return wc
    //return map[string]int{"x": 1}
}

func main() {
    //WordCount("hallo du da")
    wc.Test(WordCount)
}

