package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(10)
    go func(ch <-chan time.Time) {
        for t := range ch {
            fmt.Printf("Backup at %s\n", t)
            wg.Done()
        }
    }(time.Tick(time.Millisecond * 100))

    wg.Wait()
}