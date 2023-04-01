package series

import (
	"yiwei/database/page"
	"yiwei/database/persistence"
	"yiwei/database/series/index"
	pb "yiwei/proto"
)

func Extract(sn string) (*Series, error) {
	spb := &pb.Series{}
	if err := persistence.ExtractProto(persistence.SeriesFilePath(sn), spb); err != nil {
		return nil, err
	}

	ig, err := index.GetGenerator(spb.IndexGenerator)
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
	return persistence.DumpProto(s.spb, persistence.SeriesFilePath(s.n))
}
