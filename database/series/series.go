package series

import (
	"fmt"
	"regexp"
	"sync"

	"yiwei/database/page"
	pb "yiwei/proto"
)

var (
	nameRegex = regexp.MustCompile(`(?m)^([a-z]|[a-z][a-z0-9_\.]*[a-z0-9])$`)
)

type Series struct {
	n  string
	lp *page.Page
	rw sync.RWMutex

	spb *pb.Series
}

func Create(n string) (*Series, error) {
	if !nameRegex.MatchString(n) {
		return nil, fmt.Errorf("invalid series name: %s", n)
	}

	return &Series{
		n:  n,
		lp: page.Create(),
		spb: &pb.Series{
			IndexChain: make([]*pb.Series_IndexBlock, 0),
		},
	}, nil
}
