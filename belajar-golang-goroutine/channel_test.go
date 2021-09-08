package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Manogunawan Resqi Gultom"
		fmt.Println("Selesai Mengirim data ke Channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	//close(channel)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Manogunawan Resqi Gultom"
}

//Tes Channel Parameter
func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	//close(channel)
}

//Channel In
func OnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Manogunawan Resqi Gultom"
}

//Channel Out
func OnlyOut(channel <-chan string)  {
	data := <- channel
	fmt.Println(data)
}

//Test InOut Channel
func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

//BufferedChannel
func TestBufferedChannel(t *testing.T)  {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Manogunawan"
		channel <- "Resqi"
		channel <- "Gultom"
	}()

	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

//Range Channel
func TestRangeChannel(t *testing.T)  {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

//Select Channel
func TestSelectChannel(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for  {
		select {
		case data:= <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data:= <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
	//select {
	//case data:= <-channel1:
	//	fmt.Println("Data dari Channel 1", data)
	//case data:= <-channel2:
	//	fmt.Println("Data dari Channel 2", data)
	//}
	//
	//select {
	//case data:= <-channel1:
	//	fmt.Println("Data dari Channel 1", data)
	//case data:= <-channel2:
	//	fmt.Println("Data dari Channel 2", data)
	//}
}

//Default Select
func TestDefaultChannel(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for  {
		select {
		case data:= <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data:= <-channel2:
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