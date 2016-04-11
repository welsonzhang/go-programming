//Noempty is an example of an in-place slice algorithm.
package main

import "fmt"

//noempty returns a slice holding only the non-empty strings
//The under;ying array is modified during the call
func noempty(strings []string) []string {
    i := 0 
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}

func noempty2(strings []string) []string {
    out := strings[:0] // zero-length slice of original
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}

func remove(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
    slice[i] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}

func main() {
    data := []string{"one", "", "three"}
    fmt.Printf("%q\n", noempty2(data))
    fmt.Printf("%q\n",data)
    
    s := []int{5,6,7,8,9}
    fmt.Println(remove(s, 2))
}
