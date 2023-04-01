package page

import (
	"yiwei/data/persistence"
	pb "yiwei/proto"
)

var (
	path = persistence.GetPath("page")
)

func Extract(pid string) (*Page, error) {
	ppb := &pb.Page{}
	if err := persistence.ExtractProto(path(pid), ppb); err != nil {
		return nil, err
	}

	return &Page{id: pid, ppb: ppb}, nil
}

func (p *Page) Dump() error {
	return persistence.DumpProto(p.ppb, path(p.id))
}
