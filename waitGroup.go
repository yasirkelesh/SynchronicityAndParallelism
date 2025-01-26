package main

import (
    "fmt"
    "sync"
)
func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        fmt.Println("Hello world!")
    }()
    go func() {
        defer wg.Done()
        fmt.Println("Hello dev!")
    }()
    wg.Wait()
}