package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 1; i <= 10000; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", i))
		if err != nil {
			log.Printf("%d CLOSED (%s)\n", i, err)

			continue
		}

		log.Printf("%d OPEN\n", i)

		conn.Close()
	}

	log.Println("DONE")
}
