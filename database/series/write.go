package series

import (
	"time"

	"yiwei/database/page"
	pb "yiwei/proto"
)

func (s *Series) CommitPage() error {
	return s.lp.Commit()
}

func (s *Series) Append(val float32, ll []*pb.Label) error {
	i := time.Now().UnixNano()

	s.rw.Lock()
	defer s.rw.Unlock()

	if s.lp.IsEmpty() {
		s.spb.IndexChain = append(
			s.spb.IndexChain,
			&pb.Series_IndexBlock{
				PageId:     s.lp.ID(),
				FirstIndex: i,
			})

		if err := s.Commit(); err != nil {
			return err
		}
	}

	if err := s.lp.Append(i, val, ll); err != nil {
		return err
	}

	if s.lp.IsFull() {
		np := page.Create()
		s.lp.SetNext(np)
		return s.lp.Commit()
	}

	return nil
}
