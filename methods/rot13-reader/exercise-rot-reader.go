package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
    n, err := rr.r.Read(b)
    
    if err == io.EOF {
        return 0, err    
    }
    
    for i := 0; i < n; i++ {
        if b[i] >= 65 && b[i] <= 90 {
            b[i] = 65 + ((b[i] - 65 + 13) % 26)
        }
        if b[i] >= 97 && b[i] <= 122 {
            b[i] = 97 + ((b[i] - 97 + 13) % 26)
        }
    }
    
    return n, nil
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

