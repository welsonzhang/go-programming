package main

import (
    "fmt"
    "sort"
)

// score
type StuScore struct {
    name  string // name
    score int    // score
}

type StuScores []StuScore

//Len()
func (s StuScores) Len() int {
    return len(s)
}

//Less(): asc increase
func (s StuScores) Less(i, j int) bool {
    return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {
    stus := StuScores{
                {"alan", 95},
                {"hikerell", 91},
                {"acmfly", 96},
                {"leao", 90}}

    fmt.Println("Default:")
    for _, v := range stus {
        fmt.Println(v.name, ":", v.score)
    }
    fmt.Println()
    sort.Sort(stus)
    
    fmt.Println("Sorted:")
    for _, v := range stus {
        fmt.Println(v.name, ":", v.score)
    }

    fmt.Println("IS Sorted?", sort.IsSorted(stus))
}
