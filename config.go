package main

import "errors"

// GetPortFromConfig читает параметр PORT из конфигурации
func GetPortFromConfig(config map[string]string) (string, error) {
    port, exists := config["PORT"]
    if !exists {
        return "", errors.New("ключ PORT отсутствует в конфигурации")
    }
    return port, nil
}