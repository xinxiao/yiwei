package persistence

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	dataFileSuffix = ".dat"
)

const (
	pageDir   = "page"
	seriesDir = "series"
)

var (
	dataDir = flag.String("data_dir", "/var/lib/yiwei", "location to store yiwei data")
)

var (
	allDataDirs = make([]string, 0)
)

func PrepareDataDirectories() error {
	for _, dir := range allDataDirs {
		if err := os.MkdirAll(path.Join(*dataDir, dir), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func FilePath(dir string) func(string) string {
	allDataDirs = append(allDataDirs, dir)
	return func(f string) string {
		return path.Join(*dataDir, dir, f+dataFileSuffix)
	}
}

var (
	PageFilePath   = FilePath(pageDir)
	SeriesFilePath = FilePath(seriesDir)
)

func DirectoryScanner(dir string) func() ([]string, error) {
	return func() ([]string, error) {
		fl, err := ioutil.ReadDir(path.Join(*dataDir, dir))
		if err != nil {
			return nil, err
		}

		fnl := make([]string, len(fl))
		for i, f := range fl {
			fnl[i] = strings.TrimSuffix(f.Name(), dataFileSuffix)
		}
		return fnl, nil
	}
}

var (
	ScanSeriesDirectory = DirectoryScanner(seriesDir)
)
