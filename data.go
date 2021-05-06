package yi

import (
	"embed"
	"encoding/csv"
	"io/fs"
)

//go:embed data
var DataFiles embed.FS

func readData(f fs.File) ([][]string, error) {
	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
