package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/csaba-nagy/psps/internal/portscanner"
	"github.com/csaba-nagy/psps/internal/reporter"
)

type config struct {
	host         string
	fromPort     int
	toPort       int
	numOfWorkers int
	outputFile   string
}

type application struct {
	scanner  portscanner.PortScanner
	reporter reporter.Reporter
}

func main() {
	var cfg config
	var rp reporter.Reporter

	flag.StringVar(&cfg.host, "host", "127.0.01", "Host to scan")
	flag.IntVar(&cfg.fromPort, "from", 8080, "Port to start scanning from")
	flag.IntVar(&cfg.toPort, "to", 8090, "Port at which to stop scanning")
	flag.IntVar(&cfg.numOfWorkers, "workers", runtime.NumCPU(), "Number of the workers. Defaults to system's number of CPUs.")
	flag.StringVar(&cfg.outputFile, "output", "", "Output file path")

	flag.Parse()

	if cfg.outputFile != "" {
		rp = reporter.FileReporter{
			OutputFile: cfg.outputFile,
		}
	} else {
		rp = reporter.ConsoleReporter{}
	}

	app := &application{
		scanner:  portscanner.NewTcpPortScanner(),
		reporter: rp,
	}

	scanResult := app.scanner.Scan(portscanner.ScanQuery{
		Host:         cfg.host,
		FromPort:     cfg.fromPort,
		ToPort:       cfg.toPort,
		NumOfWorkers: cfg.numOfWorkers,
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		app.reporter.Report(scanResult.OpenPorts)

		os.Exit(0)
	}()

	err := app.reporter.Report(scanResult.OpenPorts)

	if err != nil {
		log.Fatal("Error reporting results:", err)
	}
}
