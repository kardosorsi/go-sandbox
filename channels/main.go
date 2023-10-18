package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	pages := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, page := range pages {
		go healthCheck(page, c)
	}

	//for {
		// fmt.Println(<- c) // blocking call
		// go healthCheck(<-c, c)
	// }

	for p := range c {
		go func(page string) {
			time.Sleep(5 * time.Second)
			healthCheck(page, c)
		}(p)
	}
}

func healthCheck(page string, c chan string) {
	_, err := http.Get(page) // blocking

	if err != nil {
		fmt.Println(page, " is down")
		c <- page
		return
	}

	fmt.Println(page, "is up and running")
	c <- page
}