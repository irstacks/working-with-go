// Originally from: https://gobyexample.com/channels, https://gobyexample.com/channel-buffering
// [Reprinted under license](https://creativecommons.org/licenses/by/3.0/).
// Only a few very minor fishy changes were made.

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

package main

import "fmt"

func main() {

	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
	// When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
	//
	// By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

	// Buffered channels accept a limited number of values without a corresponding receiver for those values.
	fish := []string{"one", "two", "red", "blue"}
	// Here we make a channel of strings buffering up to n values.
	multipleMessages := make(chan string, len(fish))

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	for i := range fish {
		multipleMessages <- fish[i]
	}

	for i := 0; i < len(fish)-1; i++ {
		fmt.Println(<-multipleMessages)
	}
	// But where's the blue fish?!

}
