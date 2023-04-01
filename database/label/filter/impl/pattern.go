package impl

import (
	"strings"

	"yiwei/database/label"

	"golang.org/x/exp/slices"
)

type equals struct {
	k, v string
}

func (e *equals) Approve(env *label.Env) bool {
	return env.Assert(e.k, func(v string) bool { return v == e.v })
}

type contains struct {
	k, v string
}

func (c *contains) Approve(env *label.Env) bool {
	return env.Assert(c.k, func(v string) bool { return strings.Contains(v, c.v) })
}

type in struct {
	k  string
	vl []string
}

func (i *in) Approve(env *label.Env) bool {
	return env.Assert(i.k, func(v string) bool { return slices.Contains(i.vl, v) })
}
