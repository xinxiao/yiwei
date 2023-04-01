package series

import (
	"sync"

	"yiwei/data/page"
	"yiwei/data/series/index"
	pb "yiwei/proto"
)

type Series struct {
	n  string
	ig index.Generator
	lp *page.Page

	spb *pb.Series

	rw sync.Mutex
}

func Create(n string, ign string, ctxt ...*pb.Label) (*Series, error) {
	ig, err := index.Get(ign)
	if err != nil {
		return nil, err
	}

	s := &Series{
		n:  n,
		ig: ig,
		spb: &pb.Series{
			IndexGenerator: ign,
			IndexChain:     make([]*pb.Series_IndexBlock, 0),
			ContextLabels:  ctxt,
		},
	}
	return s, nil
}

func (s *Series) Name() string {
	return s.n
}

func (s *Series) Append(val float32, ll ...*pb.Label) error {
	i := s.ig.Generate()

	if err := s.Extend(i); err != nil {
		return err
	}

	return s.lp.Append(i, val, ll...)
}

func (s *Series) Extend(i int64) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.lp != nil && !s.lp.IsFull() {
		return nil
	}

	np := page.Create()
	if s.lp == nil {
		s.lp = np
	} else {
		s.lp.SetNext(np)
		if err := s.lp.Dump(); err != nil {
			return err
		}
		s.lp = np
	}

	s.spb.IndexChain = append(
		s.spb.IndexChain,
		&pb.Series_IndexBlock{
			PageId:     s.lp.ID(),
			FirstIndex: i,
		})

	return s.Dump()
}
