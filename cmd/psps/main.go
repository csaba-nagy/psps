package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

var host string
var fromPort string
var toPort string

func init() {
	flag.StringVar(&host, "host", "127.0.01", "Host to scan")
	flag.StringVar(&fromPort, "from", "8080", "Port to start scanning from")
	flag.StringVar(&toPort, "to", "8090", "Port at which to stop scanning")
}

func main() {
	flag.Parse()

	fp, err := strconv.Atoi(fromPort)
	if err != nil {
		log.Fatalf("Invalid value at 'from': %s", fromPort)
	}

	tp, err := strconv.Atoi(toPort)
	if err != nil {
		log.Fatalf("Invalid value at 'to': %s", toPort)
	}

	if fp > tp {
		log.Fatalln("Invalid values for 'from' and 'to' port: from value cannot be higher than to")
	}

	var wg sync.WaitGroup
	goroutinesCountToWaitOn := tp - fp + 1

	wg.Add(goroutinesCountToWaitOn)

	for i := fp; i <= tp; i++ {
		go func(p int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, p))
			if err != nil {
				log.Printf("%d CLOSED (%s)\n", p, err)

				return
			}

			log.Printf("%d OPEN\n", p)

			conn.Close()
		}(i)
	}

	wg.Wait()

	log.Println("DONE")
}
