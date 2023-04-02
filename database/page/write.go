package page

import (
	pb "yiwei/proto"
)

func (p *Page) SetNext(np *Page) {
	p.ppb.NextPage = np.id
}

func (p *Page) ShouldCommit() bool {
	return false
}

func (p *Page) Append(vi int64, val float32, ll []*pb.Label) error {
	p.ppb.Entries = append(p.ppb.Entries, &pb.Entry{
		Index:  vi,
		Value:  val,
		Labels: ll,
	})

	if p.ShouldCommit() {
		return p.Commit()
	}

	return nil
}
