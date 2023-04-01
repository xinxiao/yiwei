package label

import (
	"flag"

	pb "yiwei/proto"
)

var (
	ft = flag.Int("label_flatten_threshold", 16, "label list would be flatten as map if more than the threshold")
)

type Env struct {
	l []*pb.Label
	m map[string]string
}

func CreateEnv(ll []*pb.Label) *Env {
	if len(ll) < *ft {
		return &Env{l: ll}
	}

	m := make(map[string]string)
	for _, l := range ll {
		m[l.Key] = l.Value
	}
	return &Env{m: m}
}

func (e *Env) Get(k string) (string, bool) {
	if len(e.m) > 0 {
		v, ok := e.m[k]
		return v, ok
	}

	for _, l := range e.l {
		if l.Key == k {
			return l.Value, true
		}
	}

	return "", false
}

func (e *Env) Assert(k string, asrt func(string) bool) bool {
	v, ok := e.m[k]
	return ok && asrt(v)
}
