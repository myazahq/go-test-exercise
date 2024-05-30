package main

import (
	"fmt"
	"math/rand"
	"time"
)

// producer function generates random numbers and sends them to the channel
func producer(ch chan<- int, random *rand.Rand) {
	for {
		num := random.Intn(100)
		ch <- num
		time.Sleep(time.Millisecond * 500)
	}
}

// consumer function receives numbers from the channel, calculates their square, and prints the result
func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Received %d, its square is %d\n", num, num*num)
	}
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ch := make(chan int)

	go producer(ch, r)
	go consumer(ch)

	time.Sleep(10 * time.Second)

	close(ch)
}
