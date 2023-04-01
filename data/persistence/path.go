package persistence

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var (
	dataDir = flag.String("data_dir", "/var/lib/yiwei", "location to store yiwei data")
)

var (
	allDirs = make([]string, 0)
)

func PrepareDataDirectories() error {
	for _, dir := range allDirs {
		if err := os.MkdirAll(path.Join(*dataDir, dir), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func GetPath(dir string) func(string) string {
	allDirs = append(allDirs, dir)
	return func(s string) string {
		return path.Join(*dataDir, dir, fmt.Sprintf("%s.dat", s))
	}
}
