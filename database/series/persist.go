package series

import (
	"yiwei/database/page"
	"yiwei/database/persistence"
	pb "yiwei/proto"
)

func Extract(sn string) (*Series, error) {
	spb := &pb.Series{}
	if err := persistence.ExtractProto(persistence.SeriesFilePath(sn), spb); err != nil {
		return nil, err
	}

	p, err := page.Extract(spb.IndexChain[len(spb.IndexChain)-1].PageId)
	if err != nil {
		return nil, err
	}

	return &Series{n: sn, lp: p, spb: spb}, nil
}

func (s *Series) Commit() error {
	return persistence.CommitProto(s.spb, persistence.SeriesFilePath(s.n))
}
