package page

import (
	"flag"

	dpb "yiwei/data/proto"

	"github.com/google/uuid"
)

var (
	size = flag.Uint("page_size", 256, "number of entries allowed in a single page")
)

type Page struct {
	id string
	i  int
	pb *dpb.Page
}

func Create() *Page {
	return &Page{
		id: uuid.NewString(),
		i:  0,
		pb: &dpb.Page{
			NextPage: "",
			Entries:  make([]*dpb.Page_Entry, 0, *size),
		},
	}
}

func (p *Page) ID() string {
	return p.id
}

func (p *Page) Next() string {
	return p.pb.NextPage
}

func (p *Page) IsFull() bool {
	return p.i >= cap(p.pb.Entries)
}

func (p *Page) Append(idx uint64, val float32, ll []*dpb.Label) {
	p.pb.Entries[p.i] = &dpb.Page_Entry{
		Index:  idx,
		Value:  val,
		Labels: ll,
	}
	p.i += 1
}

func (p *Page) Extend() *Page {
	np := Create()
	p.pb.NextPage = np.ID()
	return np
}
