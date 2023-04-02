package page

import (
	"flag"

	pb "yiwei/proto"
)

var (
	ac = flag.Bool("page_always_commit", false, "Commit page upon every append")
)

func (p *Page) SetNext(np *Page) {
	p.ppb.NextPage = np.id
}

func (p *Page) shouldCommit() bool {
	if *ac {
		return true
	}

	return p.Size() == 1
}

func (p *Page) Append(vi int64, val float32, ll []*pb.Label) error {
	p.ppb.Entries = append(p.ppb.Entries, &pb.Entry{
		Index:  vi,
		Value:  val,
		Labels: ll,
	})

	if p.shouldCommit() {
		return p.Commit()
	}

	return nil
}
