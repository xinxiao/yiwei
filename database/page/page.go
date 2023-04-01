package page

import (
	"flag"

	pb "yiwei/proto"

	"github.com/google/uuid"
)

var (
	size = flag.Uint("page_size", 512, "number of entries allowed in a single page")
)

type Page struct {
	id string

	ppb *pb.Page
}

func Create() *Page {
	return &Page{
		id: uuid.NewString(),
		ppb: &pb.Page{
			NextPage: "",
			Entries:  make([]*pb.Entry, 0, *size),
		},
	}
}

func (p *Page) ID() string {
	return p.id
}

func (p *Page) SetNext(np *Page) {
	p.ppb.NextPage = np.id
}

func (p *Page) IsEmpty() bool {
	return len(p.ppb.Entries) == 0
}

func (p *Page) IsFull() bool {
	return len(p.ppb.Entries) >= cap(p.ppb.Entries)
}
