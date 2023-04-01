package impl

import (
	"strings"

	"yiwei/database/label"
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
