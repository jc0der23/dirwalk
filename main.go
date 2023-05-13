package dirwalk

import (
	"os"
	"path/filepath"
	"time"
)

// LastFile returns the path of the last modified file in the directory and its subdirectories
func LastFile(directory string) (string, error) {
	var lastModifiedFile string
	var lastModificationTime time.Time

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.ModTime().After(lastModificationTime) {
			lastModifiedFile = path
			lastModificationTime = info.ModTime()
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return lastModifiedFile, nil
}
