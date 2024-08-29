package benchmarks

import (
	"os/exec"
)

func CheckWindowsFirewall() (string, error) {
	cmd := exec.Command("powershell", "Get-NetFirewallProfile | Select-Object -Property Name, Enabled")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
