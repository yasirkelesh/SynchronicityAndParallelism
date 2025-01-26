package main
import (
    "fmt"
    "time"
)

func main() {
    sayHelloWorld()
    sayHelloDev()
    time.Sleep(time.Second)
}
func sayHelloWorld() {
    fmt.Println("Hello world!")
}
func sayHelloDev() {
    fmt.Println("Hello dev!")
}