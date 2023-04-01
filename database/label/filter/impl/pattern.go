package impl

import (
	"strings"

	"yiwei/database/label"

	"golang.org/x/exp/slices"
)

type equals struct {
	k, v string
}

func (e equals) Approve(env *label.Env) bool {
	v, ok := env.Get(e.k)
	return ok && v == e.v
}

type contains struct {
	k, v string
}

func (c contains) Approve(env *label.Env) bool {
	v, ok := env.Get(c.k)
	return ok && strings.Contains(v, c.v)
}

type in struct {
	k  string
	vl []string
}

func (i in) Approve(env *label.Env) bool {
	v, ok := env.Get(i.k)
	return ok && slices.Contains(i.vl, v)
}
