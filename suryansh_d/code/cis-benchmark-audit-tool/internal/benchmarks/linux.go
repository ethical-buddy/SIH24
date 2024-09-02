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

// DisableSquashfs disables the squashfs module
func DisableSquashfs() (string, error) {
	denyCmd := exec.Command("sh", "-c", `echo "install squashfs /bin/false" >> /etc/modprobe.d/squashfs.conf && echo "blacklist squashfs" >> /etc/modprobe.d/squashfs.conf`)
	if err := denyCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to deny list squashfs: %v", err)
	}
	unloadCmd := exec.Command("sh", "-c", "modprobe -r squashfs")
	if err := unloadCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to unload squashfs: %v", err)
	}
	return "squashfs module disabled successfully.", nil
}

// DisableUdf disables the udf module
func DisableUdf() (string, error) {
	denyCmd := exec.Command("sh", "-c", `echo "install udf /bin/false" >> /etc/modprobe.d/udf.conf && echo "blacklist udf" >> /etc/modprobe.d/udf.conf`)
	if err := denyCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to deny list udf: %v", err)
	}
	unloadCmd := exec.Command("sh", "-c", "modprobe -r udf")
	if err := unloadCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to unload udf: %v", err)
	}
	return "udf module disabled successfully.", nil
}

// EnsureTmpIsSeparatePartition checks if /tmp is a separate partition
func EnsureTmpIsSeparatePartition() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /tmp | cut -d " " -f 1`)
	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check /tmp partition: %v", err)
	}
	if string(output) == "/tmp" {
		return "/tmp is already a separate partition.", nil
	}
	return "", fmt.Errorf("/tmp is not a separate partition")
}

// EnsureNodevOnTmp ensures nodev option is set on /tmp partition
func EnsureNodevOnTmp() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /tmp | grep -q "nodev"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nodev option is not set on /tmp partition")
	}
	return "nodev option is set on /tmp partition.", nil
}

// EnsureNoexecOnTmp ensures noexec option is set on /tmp partition
func EnsureNoexecOnTmp() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /tmp | grep -q "noexec"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("noexec option is not set on /tmp partition")
	}
	return "noexec option is set on /tmp partition.", nil
}

// EnsureNosuidOnTmp ensures nosuid option is set on /tmp partition
func EnsureNosuidOnTmp() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /tmp | grep -q "nosuid"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nosuid option is not set on /tmp partition")
	}
	return "nosuid option is set on /tmp partition.", nil
}

// EnsureSeparateVarPartition checks if /var is a separate partition
func EnsureSeparateVarPartition() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var | cut -d " " -f 1`)
	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check /var partition: %v", err)
	}
	if string(output) == "/var" {
		return "/var is already a separate partition.", nil
	}
	return "", fmt.Errorf("/var is not a separate partition")
}

// EnsureNodevOnVar ensures nodev option is set on /var partition
func EnsureNodevOnVar() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var | grep -q "nodev"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nodev option is not set on /var partition")
	}
	return "nodev option is set on /var partition.", nil
}

// EnsureNosuidOnVar ensures nosuid option is set on /var partition
func EnsureNosuidOnVar() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var | grep -q "nosuid"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nosuid option is not set on /var partition")
	}
	return "nosuid option is set on /var partition.", nil
}

// EnsureSeparateVarTmpPartition checks if /var/tmp is a separate partition
func EnsureSeparateVarTmpPartition() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/tmp | cut -d " " -f 1`)
	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check /var/tmp partition: %v", err)
	}
	if string(output) == "/var/tmp" {
		return "/var/tmp is already a separate partition.", nil
	}
	return "", fmt.Errorf("/var/tmp is not a separate partition")
}

// EnsureNodevOnVarTmp ensures nodev option is set on /var/tmp partition
func EnsureNodevOnVarTmp() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/tmp | grep -q "nodev"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nodev option is not set on /var/tmp partition")
	}
	return "nodev option is set on /var/tmp partition.", nil
}

// EnsureNoexecOnVarTmp ensures noexec option is set on /var/tmp partition
func EnsureNoexecOnVarTmp() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/tmp | grep -q "noexec"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("noexec option is not set on /var/tmp partition")
	}
	return "noexec option is set on /var/tmp partition.", nil
}

// EnsureNosuidOnVarTmp ensures nosuid option is set on /var/tmp partition
func EnsureNosuidOnVarTmp() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/tmp | grep -q "nosuid"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nosuid option is not set on /var/tmp partition")
	}
	return "nosuid option is set on /var/tmp partition.", nil
}

// EnsureSeparateVarLogPartition checks if /var/log is a separate partition
// func EnsureSeparateVarLogPartition() (string, error) {
// 	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log | cut -d " " -f 1`)
// 	output, err := checkCmd.CombinedOutput()
// 	if err := checkCmd.Run(); err != nil {
// 		return "", fmt.Errorf("Failed to check /var/log partition: %v", err)
// 	}
// 	if string(output) == "/var/log" {
// 		return "/var/log is already a separate partition.", nil
// 	}
// 	return "", fmt.Errorf("/var/log is not a separate partition")
// }

// // EnsureNodevOnVarLog ensures nodev option is set on /var/log partition
// func EnsureNodevOnVarLog() (string, error) {
// 	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log | grep -q "nodev"`)
// 	if err := checkCmd.Run(); err != nil {
// 		return "", fmt.Errorf("nodev option is not set on /var/log partition")
// 	}
// 	return "nodev option is set on /var/log partition.", nil
// }

// EnsureSeparateVarLogPartition checks if /var/log is a separate partition
func EnsureSeparateVarLogPartition() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log | cut -d " " -f 1`)
	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check /var/log partition: %v", err)
	}
	if string(output) == "/var/log" {
		return "/var/log is already a separate partition.", nil
	}
	return "", fmt.Errorf("/var/log is not a separate partition")
}

// EnsureNoexecOnVarLog ensures noexec option is set on /var/log partition
func EnsureNoexecOnVarLog() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log | grep -q "noexec"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("noexec option is not set on /var/log partition")
	}
	return "noexec option is set on /var/log partition.", nil
}

// EnsureNosuidOnVarLog ensures nosuid option is set on /var/log partition
func EnsureNosuidOnVarLog() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log | grep -q "nosuid"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nosuid option is not set on /var/log partition")
	}
	return "nosuid option is set on /var/log partition.", nil
}

// EnsureSeparateVarLogAuditPartition checks if /var/log/audit is a separate partition
func EnsureSeparateVarLogAuditPartition() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log/audit | cut -d " " -f 1`)
	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check /var/log/audit partition: %v", err)
	}
	if string(output) == "/var/log/audit" {
		return "/var/log/audit is already a separate partition.", nil
	}
	return "", fmt.Errorf("/var/log/audit is not a separate partition")
}

// EnsureNodevOnVarLogAudit ensures nodev option is set on /var/log/audit partition
// func EnsureNodevOnVarLogAudit() (string, error) {
// 	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log/audit | grep -q "nodev"`)
// 	if err := checkCmd.Run(); err != nil {
// 		return "", fmt.Errorf("nodev option is not set on /var/log/audit partition")
// 	}
// 	return "nodev option is set on /var/log/audit partition.", nil
// }

func EnsureNodevOnVarLog() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log | grep -q "nodev"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nodev option is not set on /var/log partition")
	}
	return "nodev option is set on /var/log partition.", nil
}

// EnsureNoexecOnVarLogAudit ensures noexec option is set on /var/log/audit partition
func EnsureNoexecOnVarLogAudit() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log/audit | grep -q "noexec"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("noexec option is not set on /var/log/audit partition")
	}
	return "noexec option is set on /var/log/audit partition.", nil
}

// EnsureNosuidOnVarLogAudit ensures nosuid option is set on /var/log/audit partition
func EnsureNosuidOnVarLogAudit() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /var/log/audit | grep -q "nosuid"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nosuid option is not set on /var/log/audit partition")
	}
	return "nosuid option is set on /var/log/audit partition.", nil
}

// EnsureSeparateHomePartition checks if /home is a separate partition
func EnsureSeparateHomePartition() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /home | cut -d " " -f 1`)
	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check /home partition: %v", err)
	}
	if string(output) == "/home" {
		return "/home is already a separate partition.", nil
	}
	return "", fmt.Errorf("/home is not a separate partition")
}

// EnsureNodevOnHome ensures nodev option is set on /home partition
func EnsureNodevOnHome() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /home | grep -q "nodev"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nodev option is not set on /home partition")
	}
	return "nodev option is set on /home partition.", nil
}

// EnsureNosuidOnHome ensures nosuid option is set on /home partition
func EnsureNosuidOnHome() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /home | grep -q "nosuid"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nosuid option is not set on /home partition")
	}
	return "nosuid option is set on /home partition.", nil
}

// EnsureNodevOnDevShm ensures nodev option is set on /dev/shm partition
func EnsureNodevOnDevShm() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /dev/shm | grep -q "nodev"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nodev option is not set on /dev/shm partition")
	}
	return "nodev option is set on /dev/shm partition.", nil
}

// EnsureNoexecOnDevShm ensures noexec option is set on /dev/shm partition
func EnsureNoexecOnDevShm() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /dev/shm | grep -q "noexec"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("noexec option is not set on /dev/shm partition")
	}
	return "noexec option is set on /dev/shm partition.", nil
}

// EnsureNosuidOnDevShm ensures nosuid option is set on /dev/shm partition
func EnsureNosuidOnDevShm() (string, error) {
	checkCmd := exec.Command("sh", "-c", `findmnt -n /dev/shm | grep -q "nosuid"`)
	if err := checkCmd.Run(); err != nil {
		return "", fmt.Errorf("nosuid option is not set on /dev/shm partition")
	}
	return "nosuid option is set on /dev/shm partition.", nil
}

// EnsureAutomountingDisabled ensures automounting is disabled
func EnsureAutomountingDisabled() (string, error) {
	checkCmd := exec.Command("sh", "-c", `systemctl is-enabled autofs`)
	output, err := checkCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check automounting status: %v", err)
	}
	if string(output) == "disabled" {
		return "Automounting is already disabled.", nil
	}
	return "", fmt.Errorf("Automounting is not disabled")
}

// EnsureUSBStorageDisabled ensures USB storage is disabled
func EnsureUSBStorageDisabled() (string, error) {
	disableCmd := exec.Command("sh", "-c", `echo "install usb-storage /bin/false" >> /etc/modprobe.d/usb-storage.conf`)
	if err := disableCmd.Run(); err != nil {
		return "", fmt.Errorf("Failed to disable USB storage: %v", err)
	}
	return "USB storage disabled successfully.", nil
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
		DisableSquashfs,
		DisableUdf,
		EnsureTmpIsSeparatePartition,
		EnsureNodevOnTmp,
		EnsureNoexecOnTmp,
		EnsureNosuidOnTmp,
		EnsureSeparateVarPartition,
		EnsureNodevOnVar,
		EnsureNosuidOnVar,
		EnsureSeparateVarTmpPartition,
		EnsureNodevOnVarTmp,
		EnsureNoexecOnVarTmp,
		EnsureNosuidOnVarTmp,
		EnsureSeparateVarLogPartition,
		//EnsureNodevOnVarLog,
		EnsureNoexecOnVarLog,
		EnsureNosuidOnVarLog,
		EnsureSeparateVarLogAuditPartition,
		//EnsureNodevOnVarLogAudit,
		EnsureNoexecOnVarLogAudit,
		EnsureNosuidOnVarLogAudit,
		EnsureSeparateHomePartition,
		EnsureNodevOnHome,
		EnsureNosuidOnHome,
		EnsureNodevOnDevShm,
		EnsureNoexecOnDevShm,
		EnsureNosuidOnDevShm,
		EnsureAutomountingDisabled,
		EnsureUSBStorageDisabled,
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

