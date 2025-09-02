package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func GetSize(p string) (string, error) {
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
				size += statI.Size()
			}
		}
	} else {
		size = stat.Size()
	}

	return fmt.Sprintf("%dB	%s", stat.Size(), stat.Name()), nil
}

func main() {
	if len(os.Args) < 1 {
		log.Fatal("No args")

		return
	}

	relativePath := os.Args[1]
	if relativePath == "" {
		log.Fatal("No file")

		return
	}

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(absPath)

	s, err := GetSize(absPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
