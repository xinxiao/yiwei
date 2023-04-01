package page

import (
	"os"
	"path"

	dpb "yiwei/data/proto"

	"google.golang.org/protobuf/proto"
)

const (
	pageDir    = "page"
	pageSuffix = ".dat"
)

func pagePath(dir string, pid string) string {
	return path.Join(dir, pageDir, pid, pageSuffix)
}

func ExtractFrom(dir string, pid string) (*Page, error) {
	src, err := os.ReadFile(pagePath(dir, pid))
	if err != nil {
		return nil, err
	}

	ppb := &dpb.Page{}
	if err := proto.Unmarshal(src, ppb); err != nil {
		return nil, err
	}

	return &Page{id: pid, i: len(ppb.Entries), pb: ppb}, nil
}

func (p *Page) DumpTo(dir string) error {
	src, err := proto.Marshal(p.pb)
	if err != nil {
		return err
	}

	return os.WriteFile(pagePath(dir, p.id), src, 0666)
}
