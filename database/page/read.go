package page

func (p *Page) Next() (*Page, error) {
	return Extract(p.ppb.NextPage)
}
