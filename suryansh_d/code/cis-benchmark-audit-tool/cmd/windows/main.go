package main

import (
    "cis-benchmark-audit-tool/internal/benchmarks"
    "cis-benchmark-audit-tool/internal/gui"
)

func main() {
    go benchmarks.RunWindowsChecks()
    gui.CreateGUI()
}

