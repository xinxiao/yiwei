package main

import (
	"flag"
	"log"

	"yiwei/data/persistence"
	"yiwei/data/series"
)

func main() {
	flag.Parse()

	if err := persistence.PrepareDataDirectories(); err != nil {
		log.Fatalf("failed to prepare data directories: %s", err)
	}

	ds, err := series.Create("test", "itoa")
	if err != nil {
		log.Fatalf("failed to create series: %s", err)
	}

	for v := float32(0.0); v <= 2048.0; v += 2.0 {
		ds.Append(v)
	}
}
