package csv

import (
	"encoding/csv"
	"os"
)

// ReadAll CSV data
func ReadAll(filepath string) ([][]string, error) {
	fl, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fl.Close()

	csvReader := csv.NewReader(fl)
	return csvReader.ReadAll()
}
