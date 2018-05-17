package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan struct{}) // 用来等待协程结束

    timer := time.NewTimer(time.Second)

    go func() {
        fmt.Printf("Now is %s\n", <-timer.C)
        done <- struct{}{}
    }()

    fmt.Println("Print in main")

    <-done
    close(done)
}