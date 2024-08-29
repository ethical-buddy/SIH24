package main

import (
    "cis-benchmark-audit-tool/internal/gui"
    "cis-benchmark-audit-tool/internal/benchmarks/linux
"
    
)

func main() {
    go benchmarks.RunLinuxChecks()
    gui.CreateGUI()
}

