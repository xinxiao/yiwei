package label

type Env struct {
	m    map[string]string
	prev *Env
}

func CreateEnv(m map[string]string) *Env {
	return &Env{m: m}
}

func (e *Env) Stack(m map[string]string) *Env {
	return &Env{m: m, prev: e}
}

func (e *Env) Get(k string) (string, bool) {
	v, ok := e.m[k]
	if !ok && e.prev != nil {
		return e.prev.Get(k)
	}
	return v, ok
}
