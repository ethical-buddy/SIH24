package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    WindowsChecks []string `json:"windows_checks"`
    LinuxChecks   []string `json:"linux_checks"`
}

func LoadConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var config Config
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}

