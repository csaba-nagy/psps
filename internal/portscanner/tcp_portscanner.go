package portscanner

import (
	"fmt"
	"net"
)

type tcpPortScanner struct{}

func (ps tcpPortScanner) Scan(query ScanQuery) ScanResult {
	result := ScanResult{}
	portsToScan := ps.makePortListToScan(query.FromPort, query.ToPort)
	portsChan := make(chan int, query.NumOfWorkers)
	resultsChan := make(chan int)

	for range cap(portsChan) {
		go ps.worker(query.Host, portsChan, resultsChan)
	}

	go func() {
		for _, p := range portsToScan {
			portsChan <- p
		}
	}()

	for range len(portsToScan) {
		if p := <-resultsChan; p != 0 {
			result.OpenPorts = append(result.OpenPorts, p)
		}
	}

	close(portsChan)
	close(resultsChan)

	return result
}

func (ps tcpPortScanner) worker(host string, portsChan <-chan int, resultsChan chan<- int) {
	for p := range portsChan {
		address := net.JoinHostPort(host, fmt.Sprintf("%d", p))
		conn, err := net.Dial("tcp", address)

		if err != nil {
			resultsChan <- 0

			continue
		}

		conn.Close()
		resultsChan <- p
	}
}

func (ps tcpPortScanner) makePortListToScan(fromPort, toPort int) []int {
	list := make([]int, 0, toPort-fromPort+1)

	for i := fromPort; i <= toPort; i++ {
		list = append(list, i)
	}

	return list
}

func NewTcpPortScanner() PortScanner {
	return tcpPortScanner{}
}
