package belajargolanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

//channel sebagai parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Chairil Ashar"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

//channel In & Out (In: simbol untuk mengirim only  <- panahnya setelah chan)
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Chairil Ashar"
}

// (Out: simbol untuk menerima only  <- panahnya sebelum chan)
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

//Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 4)
	defer close(channel)

	go func() {
		channel <- "Chairil"
		channel <- "Hekal"
		channel <- "Chairil"
		channel <- "Haifa"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Mengirim Data ke Channel")
}

//Range Channel untuk data yg banyak dan hanya 1 channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")

}

//Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

//Default Select
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}

}

//Catatan
//masukin data ke channel
// channel <- "Heril"

//ambil data dari channel
// data := <- channel
// fmt.Println(<- channel)
