package main

import "fmt"
import "strings"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
    s := make([]string, len(ip))
    for i,v := range(ip) {
        s[i] = fmt.Sprintf("%v", v)    
    }
    return fmt.Sprintf("%q", strings.Join(s, "."))
    //return fmt.Sprintf("\"%v, %v, %v, %v\"", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    ip := IPAddr{10,0,0,1}
    s := fmt.Sprintf("%v, %v, %v, %v", ip[0], ip[1], ip[2], ip[3])
    fmt.Printf("%q\n", s)
    
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}

