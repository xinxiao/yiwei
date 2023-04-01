package impl

import (
	"yiwei/database/label"
)

type and struct {
	a, b Filter
}

func (a *and) Approve(env *label.Env) bool {
	return a.a.Approve(env) && a.b.Approve(env)
}

type or struct {
	a, b Filter
}

func (o *or) Approve(env *label.Env) bool {
	return o.a.Approve(env) || o.b.Approve(env)
}

type not struct {
	f Filter
}

func (n *not) Approve(env *label.Env) bool {
	return !n.f.Approve(env)
}
