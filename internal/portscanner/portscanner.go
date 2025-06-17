package portscanner

type ScanQuery struct {
	Host         string
	FromPort     int
	ToPort       int
	NumOfWorkers int
}

type ScanResult struct {
	OpenPorts []int
}

type PortScanner interface {
	Scan(query ScanQuery) ScanResult
}
