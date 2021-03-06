package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(time.Millisecond * 1000)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 5000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
