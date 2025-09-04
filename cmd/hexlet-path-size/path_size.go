package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
)

func GetSize(p string, h bool, a bool) (string, error) {
	stat, err := os.Lstat(p)
	if err != nil {
		return "", err
	}

	var size int64 = 0
	if stat.IsDir() {
		dir, err := os.ReadDir(p)

		if err != nil {
			return "", err
		}

		for _, v := range dir {
			absPath, err := filepath.Abs(path.Join(p, v.Name()))
			if err != nil {
				return "", err
			}

			statI, err := os.Stat(absPath)
			if err != nil {
				return "", err
			}

			if statI.IsDir() == false {
				if a {
					size += statI.Size()
				} else if !strings.HasPrefix(statI.Name(), ".") {
					size += statI.Size()
				}
			}
		}
	} else {
		size = stat.Size()
	}

	sizeStr := fmt.Sprintf("%dB", size)
	if h {
		sizeStr = FormatSize(size)
	}

	return fmt.Sprintf("%s	%s", sizeStr, stat.Name()), nil
}

func FormatSize(size int64) string {
	return strings.ToUpper(strings.Replace(humanize.Bytes(uint64(size)), " ", "", -1))
}
