package series

import (
	"yiwei/data/page"
	"yiwei/data/persistence"
	"yiwei/data/series/index"
	pb "yiwei/proto"
)

var (
	path = persistence.GetPath("series")
)

func Extract(sn string) (*Series, error) {
	spb := &pb.Series{}
	if err := persistence.ExtractProto(path(sn), spb); err != nil {
		return nil, err
	}

	ig, err := index.Get(spb.IndexGenerator)
	if err != nil {
		return nil, err
	}

	s := &Series{n: sn, ig: ig, spb: spb}
	if len(spb.IndexChain) > 0 {
		p, err := page.Extract(spb.IndexChain[len(spb.IndexChain)-1].PageId)
		if err != nil {
			return nil, err
		}
		s.lp = p
	}

	return s, nil
}

func (s *Series) Dump() error {
	return persistence.DumpProto(s.spb, path(s.n))
}
