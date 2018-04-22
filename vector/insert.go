package main

import(
"fmt"
)

type X struct {
    x int
    y float64
}

const N int = 2

func main() {
    x := X{123, 2.64}
    s := make([]X, 0)
    for i := 0; i < 20; i++ {
        x := X{0, 0}
        x.x = i
        s = append(s, x)
    }

    // cut    
    i := 8
    s = append(s, x)
    copy(s[i+1:], s[i:])
    s[i] = x
	
    for i, v := range s {
        fmt.Println(i, v)
    }
}
