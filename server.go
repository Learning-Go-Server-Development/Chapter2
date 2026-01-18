package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	var msg1 string = "goroutine 1 = "
	msg2 := "goroutine 2 = "

	var ctl1 int = 10
	ctl2 := 20

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= ctl1; i++ {
			fmt.Println(msg1, i)
		}
	}()

	wg.Add(1)
	go func() {
		duration := time.Second
		time.Sleep(duration * 3)
		defer wg.Done()

		for i := 11; i <= ctl2; i++ {
			fmt.Println(msg2, i)
		}
	}()

	wg.Wait()

	port := "3000"
	msg := "Server starting on port "
	fmt.Println(msg + port)
	http.ListenAndServe(":3000", nil)
}
