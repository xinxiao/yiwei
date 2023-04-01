package series

import (
	"sync"

	"yiwei/database/page"
	"yiwei/database/series/index"
	pb "yiwei/proto"
)

type Series struct {
	n  string
	ig index.Generator
	lp *page.Page

	spb *pb.Series

	rw sync.Mutex
}

func Create(n string, igt pb.Series_IndexGeneratorType, ctxt []*pb.Label) (*Series, error) {
	ig, err := index.GetGenerator(igt)
	if err != nil {
		return nil, err
	}

	s := &Series{
		n:  n,
		ig: ig,
		spb: &pb.Series{
			IndexGenerator: igt,
			IndexChain:     make([]*pb.Series_IndexBlock, 0),
			ContextLabels:  ctxt,
		},
	}
	return s, nil
}

func (s *Series) Name() string {
	return s.n
}
