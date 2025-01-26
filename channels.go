package main

import "fmt"
func main() {
    ch := make(chan string)
    go func() {
        ch <- "Hello world!"
    }()
    go func() {
        ch <- "Hello dev!"
    }()
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}