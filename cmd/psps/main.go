package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"syscall"

	"github.com/csaba-nagy/psps/internal/portscanner"
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

	cmd := portscanner.ScanQuery{
		Host:         host,
		FromPort:     fromPort,
		ToPort:       toPort,
		NumOfWorkers: numOfWorkers,
	}

	scanner := portscanner.NewTcpPortScanner()

	scanResult := scanner.Scan(cmd)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		printResults(scanResult.OpenPorts)

		os.Exit(0)
	}()

	printResults(scanResult.OpenPorts)
}

func printResults(ports []int) {
	sort.Ints(ports)

	fmt.Println("📊 REPORT")
	fmt.Println("🚪 Opened ports")

	for _, p := range ports {
		fmt.Printf("[OPEN] %d\n", p)
	}
}
