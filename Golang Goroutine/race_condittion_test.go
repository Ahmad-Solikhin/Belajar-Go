package Golang_Goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCounterRaceCondition(t *testing.T) {
	counter := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				counter = counter + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(counter)
}
