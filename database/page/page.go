package page

import (
	"flag"
	"sync"

	pb "yiwei/proto"

	"github.com/google/uuid"
)

var (
	size = flag.Uint("page_size", 512, "number of entries allowed in a single page")
)

type Page struct {
	id string
	rw sync.RWMutex

	ppb *pb.Page
}

func Create() *Page {
	return &Page{
		id: uuid.NewString(),
		ppb: &pb.Page{
			NextPage: "",
			Entries:  make([]*pb.Entry, 0, *size),
		},
	}
}
