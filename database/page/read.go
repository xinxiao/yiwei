package page

import (
	pb "yiwei/proto"
)

func (p *Page) ID() string {
	return p.id
}

func (p *Page) Size() int {
	return len(p.ppb.Entries)
}

func (p *Page) IsEmpty() bool {
	return p.Size() == 0
}

func (p *Page) IsFull() bool {
	return p.Size() >= cap(p.ppb.Entries)
}

func (p *Page) next() string {
	return p.ppb.NextPage
}

func (p *Page) entry(i int) *pb.Entry {
	return p.ppb.Entries[i]
}

func (p *Page) Read(b, e int64, ec chan error) chan *pb.Entry {
	c := make(chan *pb.Entry)
	go p.read(b, e, c, ec)
	return c
}

func (p *Page) read(b, e int64, c chan *pb.Entry, ec chan error) {
	defer close(c)

	if p.IsEmpty() || p.entry(0).Index < b {
		return
	}

	npid := p.next()
	i, j := 0, p.Size()-1

	pfc := make(chan *pb.Entry)
	go p.prefetch(b, e, p.entry(j).Index, npid, pfc, ec)

	for i < j {
		m := (i + j) / 2
		if p.entry(m).Index < b {
			i = m + 1
		} else {
			j = m
		}
	}

	for ei := i; p.entry(ei).Index <= e; ei += 1 {
		c <- p.entry(ei)
	}

	for e := range pfc {
		c <- e
	}
}

func (p *Page) prefetch(b, e, pe int64, npid string, c chan *pb.Entry, ec chan error) {
	if pe >= e || npid == "" {
		close(c)
		return
	}

	np, err := Extract(npid)
	if err != nil {
		ec <- err
		return
	}

	go np.read(b, e, c, ec)
}
