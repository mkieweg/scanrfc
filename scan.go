package scanrfc

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

const regex = "(@rfc[[:digit:]]+|{rfc[[:digit:]]+)"

func Scan(path string) ([]string, error) {
	files, err := searchFiles(path)
	if err != nil {
		return nil, err
	}
	rfcs := make(map[string]bool)
	for _, file := range files {
		err = scanFile(file, rfcs)
		if err != nil {
			log.Error(err)
		}
	}
	out := make([]string, 0)
	for k := range rfcs {
		out = append(out, k)
	}
	return out, err
}

func searchFiles(path string) ([]string, error) {
	files := make([]string, 0)
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		filename := d.Name()
		if strings.Contains(filename, ".md") || strings.Contains(filename, ".tex") {
			file := filepath.Join(path)
			log.Debug(file)
			files = append(files, file)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("did not find any .md files")
	}
	return files, nil
}

func scanFile(path string, rfcs map[string]bool) error {
	matches := make([][]byte, 0)
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile(string(regex))
		matches = append(matches, re.FindAll([]byte(scanner.Text()), -1)...)
	}
	for _, match := range matches {
		rfcs[string(match)[1:]] = true
	}
	return nil
}
