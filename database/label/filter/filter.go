package filter

import (
	"yiwei/database/label"
	"yiwei/database/label/filter/impl"
	pb "yiwei/proto"
)

func Filter(f impl.Filter, c chan *pb.Entry) chan *pb.Entry {
	fc := make(chan *pb.Entry)
	go filter(f, c, fc)
	return fc
}

func filter(f impl.Filter, c chan *pb.Entry, fc chan *pb.Entry) {
	defer close(fc)

	for e := range c {
		if f.Approve(label.CreateEnv(e.Labels)) {
			fc <- e
		}
	}
}
