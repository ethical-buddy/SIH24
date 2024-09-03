package benchmarks

import (
	"fmt"
	"os/exec"
	"strings"
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
	if strings.TrimSpace(string(output)) == "disabled" {
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

// EnsureGPGKeysConfigured checks if GPG keys are configured
func EnsureGPGKeysConfigured() (string, error) {
	cmd := exec.Command("apt-key", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to list GPG keys: %v", err)
	}
	if len(output) == 0 {
		return "", fmt.Errorf("No GPG keys configured")
	}
	return "GPG keys are configured.", nil
}

// EnsureAppArmorInstalled checks if AppArmor is installed
func EnsureAppArmorInstalled() (string, error) {
	cmd := exec.Command("dpkg-query", "-W", "--showformat='${Status}'", "apparmor")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check AppArmor installation: %v", err)
	}
	if !strings.Contains(string(output), "install ok installed") {
		return "", fmt.Errorf("AppArmor is not installed")
	}
	return "AppArmor is installed.", nil
}

// EnsureAppArmorEnabledInBootloader checks if AppArmor is enabled in the bootloader configuration
func EnsureAppArmorEnabledInBootloader() (string, error) {
	cmd := exec.Command("grep", "'^\\s*linux'", "/boot/grub/grub.cfg")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check AppArmor in bootloader: %v", err)
	}
	if !strings.Contains(string(output), "security=apparmor") {
		return "", fmt.Errorf("AppArmor is not enabled in bootloader")
	}
	return "AppArmor is enabled in the bootloader configuration.", nil
}

// EnsureAIDEInstalled checks if AIDE is installed
func EnsureAIDEInstalled() (string, error) {
	cmd := exec.Command("dpkg-query", "-W", "--showformat='${Status}'", "aide")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check AIDE installation: %v", err)
	}
	if !strings.Contains(string(output), "install ok installed") {
		return "", fmt.Errorf("AIDE is not installed")
	}
	return "AIDE is installed.", nil
}

// EnsureUFWInstalled checks if ufw is installed
func EnsureUFWInstalled() (string, error) {
	cmd := exec.Command("dpkg-query", "-W", "--showformat='${Status}'", "ufw")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check ufw installation: %v", err)
	}
	if !strings.Contains(string(output), "install ok installed") {
		return "", fmt.Errorf("ufw is not installed")
	}
	return "ufw is installed.", nil
}

// EnsureChronyOrNTPInstalled ensures either chrony or ntp is installed
func EnsureChronyOrNTPInstalled() (string, error) {
	cmd := exec.Command("sh", "-c", "dpkg-query -W chrony ntp")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check chrony or ntp installation: %v", err)
	}
	if !strings.Contains(string(output), "chrony") && !strings.Contains(string(output), "ntp") {
		return "", fmt.Errorf("Neither chrony nor ntp is installed")
	}
	return "Chrony or NTP is installed.", nil
}

// EnsureX11ForwardingDisabled ensures X11 forwarding is disabled
func EnsureX11ForwardingDisabled() (string, error) {
	cmd := exec.Command("grep", "^X11Forwarding", "/etc/ssh/sshd_config")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check X11Forwarding setting: %v", err)
	}
	if strings.TrimSpace(string(output)) != "X11Forwarding no" {
		return "", fmt.Errorf("X11 forwarding is not disabled")
	}
	return "X11 forwarding is disabled.", nil
}

// EnsureTimeSynchronizationIsInUse ensures a time synchronization service is in use
func EnsureTimeSynchronizationIsInUse() (string, error) {
	cmd := exec.Command("timedatectl", "show", "-p", "NTPSynchronized")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check time synchronization: %v", err)
	}
	if !strings.Contains(string(output), "yes") {
		return "", fmt.Errorf("Time synchronization is not in use")
	}
	return "Time synchronization is in use.", nil
}

// EnsureNoUnnecessaryServices ensures there are no unnecessary services running
func EnsureNoUnnecessaryServices() (string, error) {
	services := []string{"avahi-daemon", "cups", "smbd", "rpcbind"}
	for _, service := range services {
		cmd := exec.Command("systemctl", "is-enabled", service)
		output, err := cmd.CombinedOutput()
		if err == nil && strings.TrimSpace(string(output)) != "disabled" {
			return "", fmt.Errorf("%s service is running", service)
		}
	}
	return "No unnecessary services are running.", nil
}

// EnsureSSHRootLoginDisabled ensures root login over SSH is disabled
func EnsureSSHRootLoginDisabled() (string, error) {
	cmd := exec.Command("grep", "^PermitRootLogin", "/etc/ssh/sshd_config")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check PermitRootLogin setting: %v", err)
	}
	if strings.TrimSpace(string(output)) != "PermitRootLogin no" {
		return "", fmt.Errorf("Root login over SSH is not disabled")
	}
	return "Root login over SSH is disabled.", nil
}

// EnsureSSHPermitEmptyPasswordsDisabled ensures SSH does not allow empty passwords
func EnsureSSHPermitEmptyPasswordsDisabled() (string, error) {
	cmd := exec.Command("grep", "^PermitEmptyPasswords", "/etc/ssh/sshd_config")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check PermitEmptyPasswords setting: %v", err)
	}
	if strings.TrimSpace(string(output)) != "PermitEmptyPasswords no" {
		return "", fmt.Errorf("SSH allows empty passwords")
	}
	return "SSH does not allow empty passwords.", nil
}

// EnsurePasswordExpirationConfigured ensures password expiration is configured
func EnsurePasswordExpirationConfigured() (string, error) {
	cmd := exec.Command("chage", "-l", "root")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check password expiration: %v", err)
	}
	if !strings.Contains(string(output), "Password expires") {
		return "", fmt.Errorf("Password expiration is not configured")
	}
	return "Password expiration is configured.", nil
}

// EnsureSSHBannerConfigured ensures the SSH banner is configured
func EnsureSSHBannerConfigured() (string, error) {
	cmd := exec.Command("grep", "^Banner", "/etc/ssh/sshd_config")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check SSH banner setting: %v", err)
	}
	if !strings.Contains(string(output), "/etc/issue.net") {
		return "", fmt.Errorf("SSH banner is not configured")
	}
	return "SSH banner is configured.", nil
}

// EnsureNISClientNotInstalled ensures the NIS client is not installed
func EnsureNISClientNotInstalled() (string, error) {
	cmd := exec.Command("dpkg-query", "-W", "--showformat='${Status}'", "nis")
	output, err := cmd.CombinedOutput()
	if err == nil && strings.Contains(string(output), "install ok installed") {
		return "", fmt.Errorf("NIS client is installed")
	}
	return "NIS client is not installed.", nil
}

// EnsureTelnetClientNotInstalled ensures the Telnet client is not installed
func EnsureTelnetClientNotInstalled() (string, error) {
	cmd := exec.Command("dpkg-query", "-W", "--showformat='${Status}'", "telnet")
	output, err := cmd.CombinedOutput()
	if err == nil && strings.Contains(string(output), "install ok installed") {
		return "", fmt.Errorf("Telnet client is installed")
	}
	return "Telnet client is not installed.", nil
}

// EnsureFTPClientNotInstalled ensures the FTP client is not installed
func EnsureFTPClientNotInstalled() (string, error) {
	cmd := exec.Command("dpkg-query", "-W", "--showformat='${Status}'", "ftp")
	output, err := cmd.CombinedOutput()
	if err == nil && strings.Contains(string(output), "install ok installed") {
		return "", fmt.Errorf("FTP client is installed")
	}
	return "FTP client is not installed.", nil
}

// EnsureIPv6IsDisabled ensures IPv6 is disabled if not in use
func EnsureIPv6IsDisabled() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv6.conf.all.disable_ipv6")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check IPv6 status: %v", err)
	}
	if !strings.Contains(string(output), "= 1") {
		return "", fmt.Errorf("IPv6 is not disabled")
	}
	return "IPv6 is disabled.", nil
}

// EnsureRootOnlyHasUID0 ensures only root has UID 0
func EnsureRootOnlyHasUID0() (string, error) {
	cmd := exec.Command("awk", "-F:", "'$3 == 0 {print $1}'", "/etc/passwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check UID 0 accounts: %v", err)
	}
	if strings.TrimSpace(string(output)) != "root" {
		return "", fmt.Errorf("More than one account has UID 0")
	}
	return "Only root has UID 0.", nil
}

// EnsureSyslogIsInstalled ensures syslog is installed
func EnsureSyslogIsInstalled() (string, error) {
	cmd := exec.Command("dpkg-query", "-W", "--showformat='${Status}'", "rsyslog")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check syslog installation: %v", err)
	}
	if !strings.Contains(string(output), "install ok installed") {
		return "", fmt.Errorf("Syslog is not installed")
	}
	return "Syslog is installed.", nil
}

// EnsureIPForwardingDisabled ensures IP forwarding is disabled
func EnsureIPForwardingDisabled() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.ip_forward")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check IP forwarding: %v", err)
	}
	if !strings.Contains(string(output), "= 0") {
		return "", fmt.Errorf("IP forwarding is not disabled")
	}
	return "IP forwarding is disabled.", nil
}

// EnsurePacketRedirectSendingDisabled ensures packet redirect sending is disabled
func EnsurePacketRedirectSendingDisabled() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.conf.all.send_redirects")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check packet redirect sending: %v", err)
	}
	if !strings.Contains(string(output), "= 0") {
		return "", fmt.Errorf("Packet redirect sending is not disabled")
	}
	return "Packet redirect sending is disabled.", nil
}

// EnsureBogusICMPResponsesIgnored ensures bogus ICMP responses are ignored
func EnsureBogusICMPResponsesIgnored() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.icmp_ignore_bogus_error_responses")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check bogus ICMP responses setting: %v", err)
	}
	if !strings.Contains(string(output), "= 1") {
		return "", fmt.Errorf("Bogus ICMP responses are not ignored")
	}
	return "Bogus ICMP responses are ignored.", nil
}

// EnsureBroadcastICMPRequestsIgnored ensures broadcast ICMP requests are ignored
func EnsureBroadcastICMPRequestsIgnored() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.icmp_echo_ignore_broadcasts")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check broadcast ICMP requests setting: %v", err)
	}
	if !strings.Contains(string(output), "= 1") {
		return "", fmt.Errorf("Broadcast ICMP requests are not ignored")
	}
	return "Broadcast ICMP requests are ignored.", nil
}

// EnsureICMPRedirectAcceptanceDisabled ensures ICMP redirects are not accepted
func EnsureICMPRedirectAcceptanceDisabled() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.conf.all.accept_redirects")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check ICMP redirect acceptance: %v", err)
	}
	if !strings.Contains(string(output), "= 0") {
		return "", fmt.Errorf("ICMP redirects are accepted")
	}
	return "ICMP redirects are not accepted.", nil
}

// EnsureSecureICMPRedirectAcceptanceDisabled ensures secure ICMP redirects are not accepted
func EnsureSecureICMPRedirectAcceptanceDisabled() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.conf.all.secure_redirects")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check secure ICMP redirect acceptance: %v", err)
	}
	if !strings.Contains(string(output), "= 0") {
		return "", fmt.Errorf("Secure ICMP redirects are accepted")
	}
	return "Secure ICMP redirects are not accepted.", nil
}

// EnsureReversePathFilteringEnabled ensures reverse path filtering is enabled
func EnsureReversePathFilteringEnabled() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.conf.all.rp_filter")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check reverse path filtering: %v", err)
	}
	if !strings.Contains(string(output), "= 1") {
		return "", fmt.Errorf("Reverse path filtering is not enabled")
	}
	return "Reverse path filtering is enabled.", nil
}

// EnsureSourceRoutedPacketsNotAccepted ensures source-routed packets are not accepted
func EnsureSourceRoutedPacketsNotAccepted() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.conf.all.accept_source_route")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check source-routed packets acceptance: %v", err)
	}
	if !strings.Contains(string(output), "= 0") {
		return "", fmt.Errorf("Source-routed packets are accepted")
	}
	return "Source-routed packets are not accepted.", nil
}

// EnsureSuspiciousPacketsLogged ensures suspicious packets are logged
func EnsureSuspiciousPacketsLogged() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.conf.all.log_martians")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check suspicious packets logging: %v", err)
	}
	if !strings.Contains(string(output), "= 1") {
		return "", fmt.Errorf("Suspicious packets are not logged")
	}
	return "Suspicious packets are logged.", nil
}

// EnsureTCPSYNCookiesEnabled ensures TCP SYN Cookies are enabled
func EnsureTCPSYNCookiesEnabled() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv4.tcp_syncookies")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check TCP SYN Cookies setting: %v", err)
	}
	if !strings.Contains(string(output), "= 1") {
		return "", fmt.Errorf("TCP SYN Cookies are not enabled")
	}
	return "TCP SYN Cookies are enabled.", nil
}

// EnsureIPv6RouterAdvertisementsNotAccepted ensures IPv6 router advertisements are not accepted
func EnsureIPv6RouterAdvertisementsNotAccepted() (string, error) {
	cmd := exec.Command("sysctl", "net.ipv6.conf.all.accept_ra")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to check IPv6 router advertisements acceptance: %v", err)
	}
	if !strings.Contains(string(output), "= 0") {
		return "", fmt.Errorf("IPv6 router advertisements are accepted")
	}
	return "IPv6 router advertisements are not accepted.", nil
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
		EnsureNoexecOnVarLog,
		EnsureNosuidOnVarLog,
		EnsureSeparateVarLogAuditPartition,
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
		EnsureGPGKeysConfigured,
		EnsureAppArmorInstalled,
		EnsureAppArmorEnabledInBootloader,
		EnsureAIDEInstalled,
		EnsureUFWInstalled,
		EnsureChronyOrNTPInstalled,
		EnsureX11ForwardingDisabled,
		EnsureTimeSynchronizationIsInUse,
		EnsureNoUnnecessaryServices,
		EnsureSSHRootLoginDisabled,
		EnsureSSHPermitEmptyPasswordsDisabled,
		EnsurePasswordExpirationConfigured,
		EnsureSSHBannerConfigured,
		EnsureNISClientNotInstalled,
		EnsureTelnetClientNotInstalled,
		EnsureFTPClientNotInstalled,
		EnsureIPv6IsDisabled,
		EnsureRootOnlyHasUID0,
		EnsureSyslogIsInstalled,
		EnsureIPForwardingDisabled,
		EnsurePacketRedirectSendingDisabled,
		EnsureBogusICMPResponsesIgnored,
		EnsureBroadcastICMPRequestsIgnored,
		EnsureICMPRedirectAcceptanceDisabled,
		EnsureSecureICMPRedirectAcceptanceDisabled,
		EnsureReversePathFilteringEnabled,
		EnsureSourceRoutedPacketsNotAccepted,
		EnsureSuspiciousPacketsLogged,
		EnsureTCPSYNCookiesEnabled,
		EnsureIPv6RouterAdvertisementsNotAccepted,
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
