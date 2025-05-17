package channel

import (
	"log"
	"math/rand"
)

const numPool = 10

func calculateValue(intCh chan int) {
	randomNumber := RandomNumber(numPool)
	intCh <- randomNumber
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

func Main() {

	// Create a channel
	// ch := make(chan int, 1)
	// ch <- 1
	// fmt.Println(<-ch)
	// fmt.Println(<-ch) // this will block the program until the value is received from the channel

	// Create a buffered channel
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// Create a channel with a goroutine
	// ch := make(chan int, 1)
	// go func() {
	// 	ch <- 1
	// }()
	// fmt.Println(<-ch)

	intCh := make(chan int)

	defer close(intCh)
	go calculateValue(intCh)
	log.Println("Random number is ", <-intCh)

}
