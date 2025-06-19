package reporter

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type FileReporter struct {
	OutputFile string
}

func (fr FileReporter) Report(result []int) error {
	f, err := os.Create(fr.OutputFile)
	if err != nil {
		return err
	}

	defer f.Close()

	date := time.Now().Format("2006-01-02 15:04:05")

	_, err = f.WriteString(fmt.Sprintf("REPORT - %s\nOPEN PORTS:\n", date))
	if err != nil {
		return err
	}

	if len(result) == 0 {
		_, err = f.WriteString("No open ports found\n")
		if err != nil {
			return err
		}

		return nil
	}

	sort.Ints(result)

	for _, p := range result {
		_, err = f.WriteString(fmt.Sprintf("%d\n", p))
		if err != nil {
			return err
		}
	}

	return nil
}
