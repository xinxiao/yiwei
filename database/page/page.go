package page

import (
	"flag"

	pb "yiwei/proto"

	"github.com/google/uuid"
)

var (
	size = flag.Int("page_size", 512, "number of entries allowed in a single page")
)

func Capacity() int {
	return *size
}

type Page struct {
	id  string
	ppb *pb.Page
}

func Create() *Page {
	return &Page{
		id: uuid.NewString(),
		ppb: &pb.Page{
			Entries: make([]*pb.Entry, 0, Capacity()),
		},
	}
}
