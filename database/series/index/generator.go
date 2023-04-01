package index

import (
	"fmt"

	pb "yiwei/proto"
)

type Generator interface {
	Generate() int64
}

func GetGenerator(igt pb.Series_IndexGeneratorType) (Generator, error) {
	switch igt {
	case pb.Series_TIMESTAMP:
		return &timestamp{}, nil
	case pb.Series_IOTA:
		return &itoa{curr: 0}, nil
	default:
		return nil, fmt.Errorf("unsupported index generator type: %s", igt)
	}
}
