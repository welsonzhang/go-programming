//Noempty is an example of an in-place slice algorithm.
package main

import "fmt"
import "github.com/garyburd/redigo/redis"

func main() {
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return 
    }
    defer c.Close()

    _, err = c.Do("SET", "mkey", "welsonzhang")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    username, err := redis.String(c.Do("GET", "mkey"))

    if err != nil {
        fmt.Println("redis get fail:", err)
    }else {
        fmt.Println("Get mkey: ", username)
    }
}
