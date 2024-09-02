package benchmarks

import (
	"fmt"
	"os/exec"
	"sync"
)

// CheckLinuxFirewall checks the status of the firewall
func CheckLinuxFirewall() (string, error) {
	cmd := exec.Command("ufw", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// DisableCramfs disables the cramfs module
func DisableCramfs() (string, error) {
	denyCmd := exec.Command("sh", "-c", `echo "install cramfs /bin/false" >> /etc/modprobe.d/cramfs.conf && echo "blacklist cramfs" >> /etc/modprobe.d/cramfs.conf`)
	if err := denyCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to deny list cramfs: %v", err)
	}
	unloadCmd := exec.Command("sh", "-c", "modprobe -r cramfs")
	if err := unloadCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to unload cramfs: %v", err)
	}
	return "cramfs module disabled successfully.", nil
}

// DisableFreevxfs disables the freevxfs module
func DisableFreevxfs() (string, error) {
	denyCmd := exec.Command("sh", "-c", `echo "install freevxfs /bin/false" >> /etc/modprobe.d/freevxfs.conf && echo "blacklist freevxfs" >> /etc/modprobe.d/freevxfs.conf`)
	if err := denyCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to deny list freevxfs: %v", err)
	}
	unloadCmd := exec.Command("sh", "-c", "modprobe -r freevxfs")
	if err := unloadCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to unload freevxfs: %v", err)
	}
	return "freevxfs module disabled successfully.", nil
}

// DisableJffs2 disables the jffs2 module
func DisableJffs2() (string, error) {
	denyCmd := exec.Command("sh", "-c", `echo "install jffs2 /bin/false" >> /etc/modprobe.d/jffs2.conf && echo "blacklist jffs2" >> /etc/modprobe.d/jffs2.conf`)
	if err := denyCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to deny list jffs2: %v", err)
	}
	unloadCmd := exec.Command("sh", "-c", "modprobe -r jffs2")
	if err := unloadCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to unload jffs2: %v", err)
	}
	return "jffs2 module disabled successfully.", nil
}

// DisableHfs disables the hfs module
func DisableHfs() (string, error) {
	denyCmd := exec.Command("sh", "-c", `echo "install hfs /bin/false" >> /etc/modprobe.d/hfs.conf && echo "blacklist hfs" >> /etc/modprobe.d/hfs.conf`)
	if err := denyCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to deny list hfs: %v", err)
	}
	unloadCmd := exec.Command("sh", "-c", "modprobe -r hfs")
	if err := unloadCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to unload hfs: %v", err)
	}
	return "hfs module disabled successfully.", nil
}

// DisableHfsplus disables the hfsplus module
func DisableHfsplus() (string, error) {
	denyCmd := exec.Command("sh", "-c", `echo "install hfsplus /bin/false" >> /etc/modprobe.d/hfsplus.conf && echo "blacklist hfsplus" >> /etc/modprobe.d/hfsplus.conf`)
	if err := denyCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to deny list hfsplus: %v", err)
	}
	unloadCmd := exec.Command("sh", "-c", "modprobe -r hfsplus")
	if err := unloadCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to unload hfsplus: %v", err)
	}
	return "hfsplus module disabled successfully.", nil
}

// RunLinuxChecks runs all the defined checks
func RunLinuxChecks() string {
	var wg sync.WaitGroup
	results := ""

	checks := []func() (string, error){
		CheckLinuxFirewall,
		DisableCramfs,
		DisableFreevxfs,
		DisableJffs2,
		DisableHfs,
		DisableHfsplus,
		// Add more Linux check functions here
	}

	for _, check := range checks {
		wg.Add(1)
		go func(chk func() (string, error)) {
			defer wg.Done()
			result, err := chk()
			if err != nil {
				results += "Error: " + err.Error() + "\n"
				return
			}
			results += result + "\n"
		}(check)
	}
	wg.Wait()

	return results
}
