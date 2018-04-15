package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1025")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "hello\n")
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
	fmt.Fprintf(conn, "Jimmy!\n")
}
