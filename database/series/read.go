package series

import (
	"yiwei/database/page"
	pb "yiwei/proto"
)

func (s *Series) Name() string {
	return s.n
}

func (s *Series) totalPageSize() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.spb.IndexChain)
}

func (s *Series) pageIndex(i int) int64 {
	return s.spb.IndexChain[i].FirstIndex
}

func (s *Series) findPage(b int64) string {
	i, j := 0, s.totalPageSize()-1
	for i < j {
		m := (i+j)/2 + 1
		if s.pageIndex(m) > b {
			j = m - 1
		} else {
			i = m
		}
	}

	return s.spb.IndexChain[i].PageId
}

func (s *Series) Read(b, e int64) (chan *pb.Entry, chan error) {
	c, ec := make(chan *pb.Entry), make(chan error)
	go s.read(b, e, c, ec)
	return c, ec
}

func (s *Series) read(b, e int64, c chan *pb.Entry, ec chan error) {
	defer close(ec)
	defer close(c)

	p, err := page.Extract(s.findPage(b))
	if err != nil {
		ec <- err
		return
	}

	for e := range p.Read(b, e, ec) {
		c <- e
	}
}
