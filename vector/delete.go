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
    //x := X{123, 2.64}
    s := make([]X, 0)
    for i := 0; i < 20; i++ {
        x := X{0, 0}
        x.x = i
        s = append(s, x)
    }

    // cut    
    i := 8
    j := 9
    copy(s[i:], s[j:])
    s = s[:len(s) - j + i]

    // delete one 
    copy(s[i:], s[i+1:])
    s = s[:len(s) - 1]
	
    for i, v := range s {
        fmt.Println(i, v)
    }
}
