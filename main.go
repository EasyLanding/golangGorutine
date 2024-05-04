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

	// timeoutFunc := timeout(3 * time.Second)
	// since := time.NewTimer(3050 * time.Millisecond)
	// for {
	// 	select {
	// 	case <-since.C:
	// 		fmt.Println("Функция не выполнена вовремя")
	// 		return
	// 	default:
	// 		if timeoutFunc() {
	// 			fmt.Println("Функция выполнена вовремя")
	// 			return
	// 		}
	// 	}
	// }

	// Создаем новый тикер с интервалом 1 секунда
	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	data := make(chan string)

	go func() {
		defer close(data)
		for {
			select {
			case <-ticker.C:
				data <- message
			case <-time.After(d):
				return
			}
		}
	}()

	return data
}