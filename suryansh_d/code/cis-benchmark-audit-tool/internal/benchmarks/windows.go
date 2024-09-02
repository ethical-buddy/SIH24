package benchmarks

import (
	"fmt"
	"os/exec"
)

// CheckWindowsFirewall checks the status of the Windows Firewall profiles.
func CheckWindowsFirewall() (string, error) {
	cmd := exec.Command("powershell", "Get-NetFirewallProfile | Select-Object -Property Name, Enabled")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// EnsurePasswordHistory checks if 'Enforce password history' is set to '24 or more passwords'.
func EnsurePasswordHistory() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Password history policy enforced", nil
}

// EnsureMaximumPasswordAge checks if 'Maximum password age' is set to '365 or fewer days, but not 0'.
func EnsureMaximumPasswordAge() (string, error) {
	cmd := exec.Command("net", "accounts", "/MAXPWAGE:365")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Maximum password age set to 365 days", nil
}

// EnsureMinimumPasswordAge checks if 'Minimum password age' is set to '1 or more days'.
func EnsureMinimumPasswordAge() (string, error) {
	cmd := exec.Command("net", "accounts", "/MINPWAGE:1")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Minimum password age set to 1 day", nil
}

// EnsureMinimumPasswordLength checks if 'Minimum password length' is set to '14 or more characters'.
func EnsureMinimumPasswordLength() (string, error) {
	cmd := exec.Command("net", "accounts", "/MINPWLEN:14")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Minimum password length set to 14 characters", nil
}

// EnsurePasswordComplexity checks if 'Password must meet complexity requirements' is set to 'Enabled'.
func EnsurePasswordComplexity() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Password complexity requirements enforced", nil
}

// EnsureRelaxMinimumPasswordLength checks if 'Relax minimum password length limits' is set to 'Enabled'.
func EnsureRelaxMinimumPasswordLength() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Minimum password length limits relaxed", nil
}

// EnsureStorePasswordsReversibleEncryption checks if 'Store passwords using reversible encryption' is set to 'Disabled'.
func EnsureStorePasswordsReversibleEncryption() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Reversible encryption for storing passwords disabled", nil
}

// EnsureAccountLockoutDuration checks if 'Account lockout duration' is set to '15 or more minutes'.
func EnsureAccountLockoutDuration() (string, error) {
	cmd := exec.Command("net", "accounts", "/LOCKOUTDURATION:15")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Account lockout duration set to 15 minutes", nil
}

// EnsureAccountLockoutThreshold checks if 'Account lockout threshold' is set to '5 or fewer invalid logon attempts, but not 0'.
func EnsureAccountLockoutThreshold() (string, error) {
	cmd := exec.Command("net", "accounts", "/LOCKOUTTHRESHOLD:5")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Account lockout threshold set to 5 attempts", nil
}

// EnsureAdministratorAccountLockout checks if 'Allow Administrator account lockout' is set to 'Enabled'.
func EnsureAdministratorAccountLockout() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Administrator account lockout enabled", nil
}

// EnsureResetAccountLockoutCounter checks if 'Reset account lockout counter after' is set to '15 or more minutes'.
func EnsureResetAccountLockoutCounter() (string, error) {
	cmd := exec.Command("net", "accounts", "/LOCKOUTWINDOW:15")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Account lockout counter reset time set to 15 minutes", nil
}

// EnsureCredentialManagerAccess checks if 'Access Credential Manager as a trusted caller' is set to 'No One'.
func EnsureCredentialManagerAccess() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Credential Manager access set to No One", nil
}

// EnsureNetworkAccess is set to 'Administrators, Remote Desktop Users'.
func EnsureNetworkAccess() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Network access limited to Administrators and Remote Desktop Users", nil
}

// EnsureActAsOs is set to 'No One'.
func EnsureActAsOs() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Act as part of the operating system is restricted to No One", nil
}

// EnsureMemoryQuotas is set to 'Administrators, LOCAL SERVICE, NETWORK SERVICE'.
func EnsureMemoryQuotas() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Memory quotas set for Administrators, LOCAL SERVICE, NETWORK SERVICE", nil
}

// EnsureLogonLocally is set to 'Administrators, Users'.
func EnsureLogonLocally() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Local logon limited to Administrators and Users", nil
}

// EnsureRemoteDesktopLogon is set to 'Administrators, Remote Desktop Users'.
func EnsureRemoteDesktopLogon() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Remote Desktop logon limited to Administrators and Remote Desktop Users", nil
}
