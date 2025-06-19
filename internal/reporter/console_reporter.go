package reporter

import (
	"fmt"
	"sort"
)

type ConsoleReporter struct{}

func (cr ConsoleReporter) Report(result []int) {
	sort.Ints(result)

	fmt.Println("📊 REPORT")

	if len(result) == 0 {
		fmt.Println("No open ports found")

		return
	}

	for _, p := range result {
		fmt.Printf("[OPEN] %d\n", p)
	}
}
