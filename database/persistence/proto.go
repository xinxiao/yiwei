package persistence

import (
	"os"

	"google.golang.org/protobuf/proto"
)

func DumpProto(pb proto.Message, p string) error {
	src, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	return os.WriteFile(p, src, 0777)
}

func ExtractProto(p string, pb proto.Message) error {
	src, err := os.ReadFile(p)
	if err != nil {
		return err
	}
	return proto.Unmarshal(src, pb)
}
