package page

func (p *Page) ID() string {
	return p.id
}

func (p *Page) Next() (*Page, error) {
	return Extract(p.ppb.NextPage)
}

func (p *Page) Size() int {
	p.rw.RLock()
	defer p.rw.RUnlock()

	return len(p.ppb.Entries)
}

func (p *Page) IsEmpty() bool {
	return p.Size() == 0
}

func (p *Page) IsFull() bool {
	return p.Size() >= cap(p.ppb.Entries)
}
