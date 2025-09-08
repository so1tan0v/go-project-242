package code

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

func GetSize(size int64, pathToObject string, all bool, recursive bool) (int64, bool, error) {
	isDir := false
	stat, err := os.Lstat(pathToObject)
	if err != nil {
		return 0, false, err
	}

	if stat.IsDir() {
		isDir = true

		dir, err := os.ReadDir(pathToObject)
		if err != nil {
			return 0, isDir, err
		}

		for _, v := range dir {
			absPath, err := filepath.Abs(path.Join(pathToObject, v.Name()))
			if err != nil {
				return 0, isDir, err
			}

			statI, err := os.Stat(absPath)
			if err != nil {
				return 0, isDir, err
			}

			if statI.IsDir() == false {
				if all {
					size += statI.Size()
				} else if !strings.HasPrefix(statI.Name(), ".") {
					size += statI.Size()
				}
			} else if recursive {
				sizeRecursive, _, err := GetSize(size, path.Join(pathToObject, statI.Name()), all, recursive)
				if err != nil {
					return 0, isDir, err
				}

				size += sizeRecursive
			}
		}
	} else {
		size += stat.Size()
	}

	return size, isDir, nil
}

func FormatSize(size int64) string {
	return strings.ToUpper(strings.Replace(humanize.Bytes(uint64(size)), " ", strconv.Itoa(0), -1))
}

func GetResult(pathToObject string, human bool, all bool, recursive bool) (string, error) {
	size, isDir, err := GetSize(0, pathToObject, all, recursive)
	if err != nil {
		log.Fatal(err)
	}

	sizeStr := fmt.Sprintf("%dB", size)
	if human {
		sizeStr = FormatSize(size)
	}

	result := fmt.Sprintf("%s\t%s", sizeStr, pathToObject)
	if isDir && pathToObject[1:] != "/" {
		result += "/"
	}

	return result, nil
}
