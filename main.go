package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a channel variable
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // execute function with goroutine
	}

	// Deprecated
	// for {
	// 	channelData := <-c // get data from channel. this code is blocking code
	// 	go checkLink(channelData, c)
	// 	go checkLink(<-c, c) // or get value directly
	// }

	// execute checkLink func with goroutine and func literal (lamda expression in c#)
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // blocking code
	if err != nil {
		fmt.Println(link + " might be down!")
		// send data into channel
		c <- link
		return
	}

	fmt.Println(link + " is up!")
	// send data into channel
	c <- link
}
