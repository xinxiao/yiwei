package impl

import (
	"yiwei/database/label"
)

type Filter interface {
	Approve(*label.Env) bool
}

func Equals(k, v string) Filter {
	return &equals{k: k, v: v}
}

func Contains(k, v string) Filter {
	return &contains{k: k, v: v}
}

func And(a, b Filter) Filter {
	return &and{a: a, b: b}
}

func Or(a, b Filter) Filter {
	return &or{a: a, b: b}
}

func Not(f Filter) Filter {
	return &not{f: f}
}
