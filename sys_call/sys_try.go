package main

import "golang.org/x/sys/unix"
import "fmt"

func main() {
   iUid := unix.Getuid()
   fmt.Println("%d", iUid)
   iTid := unix.Gettid()
   fmt.Println("%d", iTid)
}
