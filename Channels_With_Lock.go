
package main

import (
	"fmt"
	"time"
)

var c chan string
var d chan string

var lock chan bool

func order(w string, sec int) {

	time.Sleep(time.Duration(sec) * time.Second)
	c <- "great"

	time.Sleep(time.Duration(sec) * time.Second)
	d <- "good"

	fmt.Println(w, "is ready")
	lock <- true

}
func main() {
	c = make(chan string)
	d = make(chan string)
	lock = make(chan bool)

	go order("Cappucino", 4)

	go order("Milkshake", 2)

	fmt.Println("I'm waiting for my order:")

	<-c
	<-d
	<-lock
	fmt.Println(<-c)
	fmt.Println(<-d)

	time.Sleep(5 * time.Second)

}
