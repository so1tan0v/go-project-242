package code

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetSize(pathToObject string, all bool, recursive bool) (int64, error) {
	size := int64(0)
	stat, err := os.Lstat(pathToObject)
	if err != nil {
		return 0, err
	}

	if stat.IsDir() {
		dir, err := os.ReadDir(pathToObject)
		if err != nil {
			return 0, err
		}

		for _, v := range dir {
			absPath, err := filepath.Abs(path.Join(pathToObject, v.Name()))
			if err != nil {
				return 0, err
			}

			statI, err := os.Stat(absPath)
			if err != nil {
				return 0, err
			}

			if !statI.IsDir() {
				if all || !strings.HasPrefix(statI.Name(), ".") {
					size += statI.Size()
				}
			} else if recursive {
				sizeRecursive := int64(0)
				if all || !strings.HasPrefix(statI.Name(), ".") {
					sizeRecursive, err = GetSize(path.Join(pathToObject, statI.Name()), all, recursive)
					if err != nil {
						return 0, err
					}
				}

				size += sizeRecursive
			}
		}
	} else {
		size += stat.Size()
	}

	return size, nil
}

func FormatSize(fileSizeInBytes int64) string {
	double := float64(fileSizeInBytes)

	i := -1
	byteUnits := []string{"KB", "MB", "GB", "TB"}

	for double > 1024 {
		double /= 1024
		i++
	}

	if i == -1 {
		return fmt.Sprintf("%dB", fileSizeInBytes)
	}

	return fmt.Sprintf("%.1f%s", double, byteUnits[i])
}

func GetPathSize(pathToObject string, recursive, human, all bool) (string, error) {
	size, err := GetSize(pathToObject, all, recursive)
	if err != nil {
		return "", err
	}

	sizeStr := fmt.Sprintf("%dB", size)
	if human {
		sizeStr = FormatSize(size)
	}

	return sizeStr, nil
}

func GetResult(pathToObject string, human bool, all bool, recursive bool) (string, error) {
	sizeRead, err := GetPathSize(pathToObject, human, all, recursive)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s\t%s", sizeRead, pathToObject), nil
}
