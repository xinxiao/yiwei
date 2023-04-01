package index

import (
	"time"
)

type timestamp struct{}

func (*timestamp) Generate() int64 {
	return time.Now().UnixNano()
}
