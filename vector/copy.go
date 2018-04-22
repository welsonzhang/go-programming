package main

import(
"fmt"
)

type X struct {
    x int32
    y float64
}

const N int = 2

func main() {
    x := X{123, 2.64}
    s := make([]X, 0)
    for i := 0; i < N; i++ {
        s = append(s, x)
    }

    r := make([]X, len(s))
    copy(r, s)	

    for i, v := range r {
        fmt.Println(i, v)
    }
}
