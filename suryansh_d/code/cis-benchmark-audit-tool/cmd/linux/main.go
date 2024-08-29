package main

import (
	"cis-benchmark-audit-tool/internal/benchmarks"
	"cis-benchmark-audit-tool/internal/gui"
        "fmt"
		"cis-benchmark-audit-tool/internal/report"

)

func main() {

	go benchmarks.RunLinuxChecks()
    gui.CreateGUI()
    {
        r := report.Report{}
    r.AddResult("Check 1: Passed")
    r.AddResult("Check 2: Failed")
    r.AddResult("Check 3: Passed")

    err := r.GenerateReport("audit_report.txt")
    if err != nil {
        fmt.Println("Error generating report:", err)}
    }
}
