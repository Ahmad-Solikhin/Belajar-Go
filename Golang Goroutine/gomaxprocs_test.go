package Golang_Goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine", totalGoroutine)
}
