package filter

import (
	"fmt"

	"yiwei/database/label/filter/impl"
	pb "yiwei/proto"
)

func Parse(src *pb.LabelFilter) (impl.Filter, error) {
	switch op := src.Filter.(type) {
	case *pb.LabelFilter_Equals_:
		return impl.Equals(op.Equals.Key, op.Equals.Value), nil
	case *pb.LabelFilter_Contains_:
		return impl.Contains(op.Contains.Key, op.Contains.Value), nil
	case *pb.LabelFilter_Not_:
		f, err := Parse(op.Not.Base)
		if err != nil {
			return nil, err
		}

		return impl.Not(f), nil
	case *pb.LabelFilter_And_:
		a, err := Parse(op.And.First)
		if err != nil {
			return nil, err
		}

		b, err := Parse(op.And.Second)
		if err != nil {
			return nil, err
		}

		return impl.And(a, b), nil
	case *pb.LabelFilter_Or_:
		a, err := Parse(op.Or.First)
		if err != nil {
			return nil, err
		}

		b, err := Parse(op.Or.Second)
		if err != nil {
			return nil, err
		}

		return impl.Or(a, b), nil
	default:
		return nil, fmt.Errorf("unsupported type of filter: %T", op)
	}
}
