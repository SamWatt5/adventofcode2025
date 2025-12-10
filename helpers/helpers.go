package helpers

import (
	"os"
)

// readFile reads the entire file and returns its content as a string
func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Check(err error) {
	if err != nil {
		panic(err)
	}

}
