package label

import (
	"fmt"

	pb "yiwei/proto"
)

func AsMap(ll []*pb.Label) (map[string]string, error) {
	cm := make(map[string]string)
	for _, l := range ll {
		if _, ok := cm[l.Key]; ok {
			return nil, fmt.Errorf("duplicated key in context: %s", l.Key)
		} else {
			cm[l.Key] = l.Value
		}
	}
	return cm, nil
}
