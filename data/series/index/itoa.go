package index

type itoa struct {
	curr int64
}

func (*itoa) Name() string {
	return "itoa"
}

func (ii *itoa) Generate() int64 {
	curr := ii.curr
	ii.curr += 1
	return curr
}
