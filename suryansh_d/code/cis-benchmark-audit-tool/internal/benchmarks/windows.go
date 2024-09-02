package benchmarks

import (
	//"fmt"
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

// EnsureCreatePagefile checks if 'Create a pagefile' is set to 'Administrators'.
func EnsureCreatePagefile() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Create a pagefile set to Administrators", nil
}

// EnsureCreateTokenObject checks if 'Create a token object' is set to 'No One'.
func EnsureCreateTokenObject() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Create a token object set to No One", nil
}

// EnsureCreateGlobalObjects checks if 'Create global objects' is set to 'Administrators, LOCAL SERVICE, NETWORK SERVICE, SERVICE'.
func EnsureCreateGlobalObjects() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Create global objects set to Administrators, LOCAL SERVICE, NETWORK SERVICE, SERVICE", nil
}

// EnsureDenyLogonLocally checks if 'Deny log on locally' is set to include 'Guests'.
func EnsureDenyLogonLocally() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Deny log on locally set to include Guests", nil
}

// EnsureDenyLogonThroughRemoteDesktop checks if 'Deny log on through Remote Desktop Services' is set to include 'Guests, Local account'.
func EnsureDenyLogonThroughRemoteDesktop() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Deny log on through Remote Desktop Services set to include Guests, Local account", nil
}

// EnsureForceShutdownFromRemoteSystem checks if 'Force shutdown from a remote system' is set to 'Administrators'.
func EnsureForceShutdownFromRemoteSystem() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Force shutdown from a remote system set to Administrators", nil
}

// EnsureGenerateSecurityAudits checks if 'Generate security audits' is set to 'LOCAL SERVICE, NETWORK SERVICE'.
func EnsureGenerateSecurityAudits() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Generate security audits set to LOCAL SERVICE, NETWORK SERVICE", nil
}

// EnsureIncreaseSchedulingPriority checks if 'Increase scheduling priority' is set to 'Administrators, Window Manager\\Window Manager Group'.
func EnsureIncreaseSchedulingPriority() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Increase scheduling priority set to Administrators, Window Manager\\Window Manager Group", nil
}

// EnsureLoadUnloadDeviceDrivers checks if 'Load and unload device drivers' is set to 'Administrators'.
func EnsureLoadUnloadDeviceDrivers() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Load and unload device drivers set to Administrators", nil
}

// EnsureAccountsBlockMicrosoftAccounts checks if 'Accounts: Block Microsoft accounts' is set to 'Users can't add or log on with Microsoft accounts'.
func EnsureAccountsBlockMicrosoftAccounts() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Accounts: Block Microsoft accounts set to 'Users can't add or log on with Microsoft accounts'", nil
}

// EnsureGuestAccountStatus checks if 'Accounts: Guest account status' is set to 'Disabled'.
func EnsureGuestAccountStatus() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Accounts: Guest account status set to Disabled", nil
}

// EnsureLimitBlankPasswordUsage checks if 'Accounts: Limit local account use of blank passwords to console logon only' is set to 'Enabled'.
func EnsureLimitBlankPasswordUsage() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Accounts: Limit local account use of blank passwords to console logon only set to Enabled", nil
}

// EnsureAuditForceAuditPolicySubcategorySettings checks if 'Audit: Force audit policy subcategory settings (Windows Vista or later) to override audit policy category settings' is set to 'Enabled'.
func EnsureAuditForceAuditPolicySubcategorySettings() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Audit: Force audit policy subcategory settings (Windows Vista or later) to override audit policy category settings set to Enabled", nil
}

// EnsureAuditShutDownSystemIfUnableToLogSecurityAudits checks if 'Audit: Shut down system immediately if unable to log security audits' is set to 'Disabled'.
func EnsureAuditShutDownSystemIfUnableToLogSecurityAudits() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Audit: Shut down system immediately if unable to log security audits set to Disabled", nil
}

// EnsureDCOMRestrictions checks if 'DCOM: Machine Access Restrictions in Security Descriptor Definition Language (SDDL) syntax' is set to 'Enabled'.
func EnsureDCOMRestrictions() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "DCOM: Machine Access Restrictions set to Enabled", nil
}

// EnsureDomainMemberDisableMachineAccountPasswordChanges checks if 'Domain member: Disable machine account password changes' is set to 'Disabled'.
func EnsureDomainMemberDisableMachineAccountPasswordChanges() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Domain member: Disable machine account password changes set to Disabled", nil
}

// EnsureDomainMemberMaxMachineAccountPasswordAge checks if 'Domain member: Maximum machine account password age' is set to '30 or fewer days, but not 0'.
func EnsureDomainMemberMaxMachineAccountPasswordAge() (string, error) {
	cmd := exec.Command("net", "accounts", "/MAXPWAGE:30")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Domain member: Maximum machine account password age set to 30 or fewer days", nil
}

// EnsureDomainMemberRequireStrongSessionKey checks if 'Domain member: Require strong (Windows 2000 or later) session key' is set to 'Enabled'.
func EnsureDomainMemberRequireStrongSessionKey() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Domain member: Require strong (Windows 2000 or later) session key set to Enabled", nil
}

// EnsureInteractiveLogonDoNotRequireCtrlAltDel checks if 'Interactive logon: Do not require CTRL+ALT+DEL' is set to 'Disabled'.
func EnsureInteractiveLogonDoNotRequireCtrlAltDel() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Interactive logon: Do not require CTRL+ALT+DEL set to Disabled", nil
}

// EnsureInteractiveLogonDontDisplayLastSignedIn checks if 'Interactive logon: Don't display last signed-in' is set to 'Enabled'.
func EnsureInteractiveLogonDontDisplayLastSignedIn() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Interactive logon: Don't display last signed-in set to Enabled", nil
}

// EnsureInteractiveLogonMachineAccountLockoutThreshold checks if 'Interactive logon: Machine account lockout threshold' is set to '10 or fewer invalid logon attempts, but not 0'.
func EnsureInteractiveLogonMachineAccountLockoutThreshold() (string, error) {
	cmd := exec.Command("net", "accounts", "/LOCKOUTTHRESHOLD:10")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Interactive logon: Machine account lockout threshold set to 10 or fewer invalid logon attempts", nil
}

// EnsureInteractiveLogonMachineInactivityLimit checks if 'Interactive logon: Machine inactivity limit' is set to '900 or fewer seconds, but not 0'.
func EnsureInteractiveLogonMachineInactivityLimit() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Interactive logon: Machine inactivity limit set to 900 or fewer seconds", nil
}

// EnsureNetworkAccessLetEveryonePermissionsApplyToAnonymousUsers checks if 'Network access: Let Everyone permissions apply to anonymous users' is set to 'Disabled'.
func EnsureNetworkAccessLetEveryonePermissionsApplyToAnonymousUsers() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Network access: Let Everyone permissions apply to anonymous users set to Disabled", nil
}

// EnsureNetworkAccessDoNotAllowAnonymousEnumOfSAMAccountsAndShares checks if 'Network access: Do not allow anonymous enumeration of SAM accounts and shares' is set to 'Enabled'.
func EnsureNetworkAccessDoNotAllowAnonymousEnumOfSAMAccountsAndShares() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Network access: Do not allow anonymous enumeration of SAM accounts and shares set to Enabled", nil
}

// EnsureNetworkAccessRestrictAnonymousAccessToNamedPipesAndShares checks if 'Network access: Restrict anonymous access to Named Pipes and Shares' is set to 'Enabled'.
func EnsureNetworkAccessRestrictAnonymousAccessToNamedPipesAndShares() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Network access: Restrict anonymous access to Named Pipes and Shares set to Enabled", nil
}

// EnsureNetworkSecurityLANManagerAuthenticationLevel checks if 'Network security: LAN Manager authentication level' is set to 'Send NTLMv2 response only. Refuse LM & NTLM'.
func EnsureNetworkSecurityLANManagerAuthenticationLevel() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Network security: LAN Manager authentication level set to 'Send NTLMv2 response only. Refuse LM & NTLM'", nil
}

// EnsureNetworkSecurityDoNotStoreLANManagerHashValue checks if 'Network security: Do not store LAN Manager hash value on next password change' is set to 'Enabled'.
func EnsureNetworkSecurityDoNotStoreLANManagerHashValue() (string, error) {
	cmd := exec.Command("secedit", "/export", "/cfg", "secpol.cfg")
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "Network security: Do not store LAN Manager hash value on next password change set to Enabled", nil
}
