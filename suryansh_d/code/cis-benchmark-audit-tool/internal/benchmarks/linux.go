package benchmarks

import (
    "os/exec"
    "sync"
    "fmt"
)

func CheckLinuxFirewall() (string, error) {
    cmd := exec.Command("ufw", "status")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(output), nil
}

func RunLinuxChecks() {
    var wg sync.WaitGroup
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
                return
            }
            // Use the result variable here, for example:
            fmt.Println(result)
        }(check)
    }
    wg.Wait()
}

