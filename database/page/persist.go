package page

import (
	"yiwei/database/persistence"
	pb "yiwei/proto"
)

func Extract(pid string) (*Page, error) {
	ppb := &pb.Page{}
	if err := persistence.ExtractProto(persistence.PageFilePath(pid), ppb); err != nil {
		return nil, err
	}

	return &Page{id: pid, ppb: ppb}, nil
}

func (p *Page) Dump() error {
	return persistence.DumpProto(p.ppb, persistence.PageFilePath(p.id))
}
