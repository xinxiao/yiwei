package page

import (
	"yiwei/database/label"
	pb "yiwei/proto"
)

func (p *Page) SetNext(np *Page) {
	p.ppb.NextPage = np.id
}

func (p *Page) Append(vi int64, val float32, ll []*pb.Label) error {
	e := &pb.Entry{
		Index: vi,
		Value: val,
	}

	if len(ll) > 0 {
		lm, err := label.AsMap(ll)
		if err != nil {
			return err
		}
		e.Labels = lm
	}

	p.ppb.Entries = append(p.ppb.Entries, e)
	return nil
}
