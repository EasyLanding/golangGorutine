package main

import (
	"fmt"
	"time"
)

func HelloWorld() string {
	return "Hello world!"
}

func timeout(timeout time.Duration) func() bool {
	ch := make(chan struct{})

	go func() {
		time.Sleep(timeout)
		close(ch)
	}()

	return func() bool {
		select {
		case <-ch:
			return false
		default:
			return true
		}
	}
}

func main() {
	helloWorld := HelloWorld()
	fmt.Println(helloWorld)

	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)
	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}
