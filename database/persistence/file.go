package persistence

import (
	"os"
)

func WriteToFile(b []byte, p string) error {
	return os.WriteFile(p, b, 0777)
}

func ReadFromFile(p string) ([]byte, error) {
	return os.ReadFile(p)
}
