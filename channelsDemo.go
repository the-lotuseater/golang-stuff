package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("A channel is way for bidirectional communication between go routines.")
	fmt.Println("There are 2 types of channels, UnBuffered (default) and buffered.")

	fmt.Println("A UnBuffered channel takes a receiver and sender. The abscence of a receiver results in a deadlock.")

	fmt.Println("A Buffered Channel has a buffer that can be used as a way to store data (a buffer) until a receiver joins the Channel. This is sort of async communication.")
	fmt.Println("Create a bidirectional Channel")
	ch1 := make(chan string)
	go getName(ch1)

	name := <-ch1
	fmt.Println("Received output from the channel. Printing username")
	fmt.Println("Hello ", name, "!")

	chan2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6}
	var multiplier int = 10
	fmt.Println("Before multiplying array by: ", multiplier)
	go multiplyArray(chan2, arr, multiplier)
	divisor := <-chan2
	fmt.Println("After multiplyArray", divisor)
	go divideArray(chan2, divisor)

	finalResult := <-chan2
	fmt.Println("Final after dividing by:", finalResult)

	fmt.Println("##########")
	fmt.Println("Testing sum")
	chan3 := make(chan int)
	go sum(chan3, arr[:3])
	go sum(chan3, arr[3:])

	sum1, sum2 := <-chan3, <-chan3
	fmt.Println("Sum1 and Sum2 are. This demonstrates that the Channels are LIFO ", sum1, sum2)

}

func getName(channel chan string) {
	fmt.Println("Adding username into the channel. As you can see the main thread will stay locked until it receives a response from this channel. This allows easy synchronization between routines.")
	time.Sleep(10 * time.Second)
	channel <- "Abhishek"
}

func multiplyArray(channel chan int, arr []int, weight int) {
	var product int = 1
	for i := 0; i < len(arr); i++ {
		product *= arr[i] * weight
	}
	fmt.Println("Sleeping thread for 10 seconds, this will block the main thread while it waits for the result of the operation to complete.")
	time.Sleep(weight * time.Second)
	channel <- product
}

func divideArray(channel chan int, divisor int) {
	var res int = divisor / 2
	channel <- res
}

func sum(channel chan int, arr []int) {
	var sum int = 0
	for idx := 0; idx < len(arr); idx++ {
		sum += arr[idx]
	}
	channel <- sum
}

