package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/csaba-nagy/psps/internal/portscanner"
	"github.com/csaba-nagy/psps/internal/reporter"
)

var (
	host         string
	fromPort     int
	toPort       int
	numOfWorkers int
)

type application struct {
	scanner  portscanner.PortScanner
	reporter reporter.Reporter
}

func init() {
	flag.StringVar(&host, "host", "127.0.01", "Host to scan")
	flag.IntVar(&fromPort, "from", 8080, "Port to start scanning from")
	flag.IntVar(&toPort, "to", 8090, "Port at which to stop scanning")
	flag.IntVar(&numOfWorkers, "workers", runtime.NumCPU(), "Number of the workers. Defaults to system's number of CPUs.")
}

func main() {
	flag.Parse()

	application := &application{
		scanner:  portscanner.NewTcpPortScanner(),
		reporter: reporter.ConsoleReporter{},
	}

	scanResult := application.scanner.Scan(portscanner.ScanQuery{
		Host:         host,
		FromPort:     fromPort,
		ToPort:       toPort,
		NumOfWorkers: numOfWorkers,
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		application.reporter.Report(scanResult.OpenPorts)

		os.Exit(0)
	}()

	application.reporter.Report(scanResult.OpenPorts)
}
