package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"syscall"
)

var (
	host         string
	fromPort     int
	toPort       int
	numOfWorkers int
)

func init() {
	flag.StringVar(&host, "host", "127.0.01", "Host to scan")
	flag.IntVar(&fromPort, "from", 8080, "Port to start scanning from")
	flag.IntVar(&toPort, "to", 8090, "Port at which to stop scanning")
	flag.IntVar(&numOfWorkers, "workers", runtime.NumCPU(), "Number of the workers. Defaults to system's number of CPUs.")
}

func main() {
	flag.Parse()

	var openPorts []int

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		printResults(openPorts)

		os.Exit(0)
	}()

	portsToScan := getPortListToScan(fromPort, toPort)

	portsChan := make(chan int, numOfWorkers)
	resultsChan := make(chan int)

	for range cap(portsChan) {
		go worker(host, portsChan, resultsChan)
	}

	go func() {
		for _, p := range portsToScan {
			portsChan <- p
		}
	}()

	for range len(portsToScan) {
		if p := <-resultsChan; p != 0 {
			openPorts = append(openPorts, p)
		}
	}

	close(portsChan)
	close(resultsChan)

	printResults(openPorts)
}

func worker(host string, portsChan <-chan int, resultsChan chan<- int) {
	for p := range portsChan {
		address := fmt.Sprintf("%s:%d", host, p)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			resultsChan <- 0

			continue
		}

		conn.Close()
		resultsChan <- p
	}
}

func printResults(ports []int) {
	sort.Ints(ports)

	fmt.Println("📊 REPORT")
	fmt.Println("🚪Opened ports")

	for _, p := range ports {
		fmt.Printf("%d\n", p)
	}
}

func getPortListToScan(fromPort, toPort int) []int {
	list := make([]int, 0, toPort-fromPort+1)

	for i := fromPort; i <= toPort; i++ {
		list = append(list, i)
	}

	return list
}
