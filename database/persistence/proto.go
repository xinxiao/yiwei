package persistence

import (
	"google.golang.org/protobuf/proto"
)

func CommitProto(pb proto.Message, p string) error {
	b, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	return WriteToFile(b, p)
}

func ExtractProto(p string, pb proto.Message) error {
	b, err := ReadFromFile(p)
	if err != nil {
		return err
	}

	return proto.Unmarshal(b, pb)
}
