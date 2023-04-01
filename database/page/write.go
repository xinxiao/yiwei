package page

import (
	pb "yiwei/proto"
)

func (p *Page) Append(vi int64, val float32, ll []*pb.Label) error {
	p.ppb.Entries = append(p.ppb.Entries, &pb.Entry{
		Index:  vi,
		Value:  val,
		Labels: ll,
	})
	return nil
}
