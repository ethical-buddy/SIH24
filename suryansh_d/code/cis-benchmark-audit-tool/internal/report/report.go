package report

import (
    "os"
    "fmt"
)

type Report struct {
    Results []string
}

func (r *Report) AddResult(result string) {
    r.Results = append(r.Results, result)
}

func (r *Report) GenerateReport(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    for _, result := range r.Results {
        _, err := file.WriteString(result + "\n")
        if err != nil {
            return err
        }
    }
    fmt.Println("Report generated:", filename)
    return nil
}

