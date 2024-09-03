package main

import (
	"log"
	"os/exec"
)

func installPackages(packages []string) {
	for _, pkg := range packages {
		cmd := exec.Command("sudo", "apt-get", "install", "-y", pkg)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to install package %s: %v", pkg, err)
		}
		log.Printf("Successfully installed package: %s", pkg)
	}
}

func main() {
	// List of packages required for the benchmarks
	packages := []string{
		"iptables",          // For firewall checks
		"e2fsprogs",         // For filesystems checks
		"lvm2",              // For partition checks
		"usbutils",          // For USB storage checks
		"automount",         // For automounting checks
		"dosfstools",        // For filesystem checks
		"bsdmainutils",      // For additional system utilities
		"systemd-sysv",      // For systemd and service management
	}

	installPackages(packages)
}

