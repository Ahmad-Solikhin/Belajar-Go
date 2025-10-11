package Golang_Goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func SendDataToChannel(channel chan string) {
	fmt.Println("Start sleep")
	time.Sleep(2 * time.Second)

	channel <- "Gayuh"
	fmt.Println("Done send data")
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go SendDataToChannel(channel)

	fmt.Println(<-channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Ahmad"
}

func OnlyOut(channel <-chan string) {
	fmt.Println(<-channel)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyOut(channel)
	go OnlyIn(channel)

	time.Sleep(1 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Gayuh"
	channel <- "Raharjo"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Display " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		time.Sleep(1 * time.Second)
		fmt.Println("Menerima data", data)
	}

}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go SendDataToChannel(channel1)
	go SendDataToChannel(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
