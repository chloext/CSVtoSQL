package csv

import (
	"encoding/csv"
	"errors"
	"os"
)

func CSVReader(paths []string) (files [][][]string, err error) {
	for _, p := range paths {
		file, err := os.Open(p)
		if err != nil {
			return nil, errors.New("invalid file input" + p)
		}
		defer file.Close()

		r := csv.NewReader(file)
		f, err := r.ReadAll()
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return
}
