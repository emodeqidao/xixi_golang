package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string)
	for i := 0; i <5 ; i++  {
		go printhelloword(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println("chan:"+ msg)
	}

	time.Sleep(time.Millisecond)
}

func printhelloword(i int,  ch chan string)  {

		ch <- fmt.Sprintf("hello word %d\n", i)

}
