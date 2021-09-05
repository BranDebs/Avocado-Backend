// Package secrets handles interfacing with docker secrets.
package secrets

import (
	"bufio"
	"fmt"
	"os"
)

// SingleLineKey returns the key from a single line in a secrets file.
func SingleLineKey(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("secrets.SingleLineKey open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, fmt.Errorf("secrets.SingleLineKey scan file content: %w", scanner.Err())
	}

	return scanner.Bytes(), nil
}
