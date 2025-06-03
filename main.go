package main

import "time"

func channelOwner[T any](data T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- data
		}
	}()
	return ch
}

func channelConsumer[T any](ch <-chan T) {
	for data := range ch {
		println(data)
	}
}

func main() {

	ch := channelOwner[string]("Hello, World!")

	channelConsumer(ch)

}
