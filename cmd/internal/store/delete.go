package store

import (
	"fmt"
	"os"
	"path/filepath"
)

func Delete(name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("invalid name")
	}
	path := filepath.Join(home, ".apictl", "requests", name+".json")
	
	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("invalid name")
	}

	return nil
}