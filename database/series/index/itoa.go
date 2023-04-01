package index

type itoa struct {
	curr int64
}

func (ii *itoa) Generate() int64 {
	curr := ii.curr
	ii.curr += 1
	return curr
}
