package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestHelloWorldFunc(t *testing.T) {
	var buf bytes.Buffer

	fmt.Fprint(&buf, HelloWorld())

	output := buf.String()
	expected := "Hello world!"

	if output != expected {
		t.Errorf("Unexpected output: %s", output)
	}
}

func TestTimeout(t *testing.T) {
	timeoutFunc := timeout(3 * time.Second)

	// Проверка выполнения функции вовремя
	startTime := time.Now()
	if !timeoutFunc() {
		t.Errorf("Функция должна была выполниться вовремя")
	}
	elapsedTime := time.Since(startTime)
	if elapsedTime > 1*time.Second {
		t.Errorf("Функция выполнилась слишком долго: %v", elapsedTime)
	}

	// Проверка невыполнения функции вовремя
	time.Sleep(4 * time.Second)
	if timeoutFunc() {
		t.Errorf("Функция должна была завершиться по таймауту")
	}
}
