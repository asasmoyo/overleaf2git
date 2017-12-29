package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func AddFiles(src, dest string) {
	dstFiles, err := ioutil.ReadDir(dest)
	if err != nil {
		panic(err)
	}
	for _, item := range dstFiles {
		if strings.TrimSpace(item.Name()) != ".git" {
			path := filepath.Join(dest, item.Name())
			os.RemoveAll(path)
		}
	}

	if !strings.HasSuffix(src, "/") {
		src += "/."
	}
	if !strings.HasSuffix(dest, "/") {
		dest += "/"
	}
	run("cp", "-a", src, dest)
}
