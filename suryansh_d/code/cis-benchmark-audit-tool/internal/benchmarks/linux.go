package benchmarks

import (
	"os/exec"
	"sync"
       )

func CheckLinuxFirewall() (string, error) {
	cmd := exec.Command("ufw", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func RunLinuxChecks() string {
	var wg sync.WaitGroup
	results := ""

	checks := []func() (string, error){
		CheckLinuxFirewall,
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

