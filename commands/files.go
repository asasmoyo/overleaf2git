package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func AddFiles(src, dest string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		panic(err)
	}

	for _, item := range files {
		if !(item.IsDir() && item.Name() != ".git") {
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
