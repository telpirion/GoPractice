package main

import "fmt"

func main() {
	passMessagesBetweenGoroutines("Hello there")
}

func passMessagesBetweenGoroutines(message string) {
	c1 := make(chan int)
	c2 := make(chan int)
	p := make(chan string)

	max := len(message)

	go printCharAtIndex(message, "c1", max, c1, c2, p)
	go printCharAtIndex(message, "c2", max, c2, c1, p)

	c1 <- 0

	for s := range p {
		fmt.Println(s)
	}
}

func printCharAtIndex(message, chName string, maxLength int, c chan int, o chan int, p chan string) {

	for {
		if i, ok := <-c; ok {
			if i < maxLength {
				chr := message[i]

				p <- fmt.Sprintf("%s: %s", chName, string(chr))
				o <- i + 1
			} else {
				break
			}
		}
	}

	close(c)
	close(p)
}
