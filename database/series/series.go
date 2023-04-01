package series

import (
	"fmt"
	"regexp"
	"sync"

	"yiwei/database/label"
	"yiwei/database/page"
	"yiwei/database/series/index"
	pb "yiwei/proto"
)

var (
	nameRegex = regexp.MustCompile(`(?m)^([a-z]|[a-z][a-z0-9_\.]*[a-z0-9])$`)
)

type Series struct {
	n  string
	ig index.Generator
	le *label.Env
	lp *page.Page
	rw sync.RWMutex

	spb *pb.Series
}

func Create(n string, igt pb.Series_IndexGeneratorType, ll []*pb.Label) (*Series, error) {
	if !nameRegex.MatchString(n) {
		return nil, fmt.Errorf("invalid series name: %s", n)
	}

	ig, err := index.GetGenerator(igt)
	if err != nil {
		return nil, err
	}

	cm, err := label.AsMap(ll)
	if err != nil {
		return nil, err
	}

	return &Series{
		n:  n,
		ig: ig,
		le: label.CreateEnv(cm),
		lp: page.Create(),
		spb: &pb.Series{
			IndexGenerator: igt,
			IndexChain:     make([]*pb.Series_IndexBlock, 0),
			ContextLabels:  ll,
		},
	}, nil
}
