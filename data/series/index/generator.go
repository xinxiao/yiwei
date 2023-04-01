package index

import (
	"fmt"
)

type Generator interface {
	Name() string
	Generate() int64
}

var (
	igl = []Generator{
		&itoa{curr: 0},
		&timestamp{},
	}
)

func Get(n string) (Generator, error) {
	for _, ig := range igl {
		if ig.Name() == n {
			return ig, nil
		}
	}
	return nil, fmt.Errorf("unsupported index generator: %s", n)
}
