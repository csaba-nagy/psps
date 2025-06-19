package reporter

import (
	"fmt"
	"sort"
)

type ConsoleReporter struct{}

func (cr ConsoleReporter) Report(result []int) error {
	sort.Ints(result)

	fmt.Println("📊 REPORT")

	if len(result) == 0 {
		fmt.Println("No open ports found")

		return nil
	}

	for _, p := range result {
		fmt.Printf("[OPEN] %d\n", p)
	}

	return nil
}
