package datapreparation

import (
	"bufio"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/pkg/errors"
)

func LoadCSV(file *os.File) (*dataframe.DataFrame, error) {
	r := bufio.NewReader(file)

	df := dataframe.ReadCSV(r)

	if err := df.Err; err != nil {
		return nil, errors.Wrap(err, "error reading CSV")
	}

	if len(df.Names()) == 0 {
		return nil, errors.New("CSV file has no header row")
	}

	return &df, nil
}

func DropCols(df dataframe.DataFrame, colNames []string) (dataframe.DataFrame, error) {
	if len(colNames) == 0 {
		return dataframe.DataFrame{}, errors.New("at least one column name must be provided")
	}

	// Create a map of column names to drop for quick lookup
	dropCols := make(map[string]bool)
	for _, colName := range colNames {
		dropCols[colName] = true
	}

	// Create a list of column names to keep
	var keepCols []string
	for _, colName := range df.Names() {
		if !dropCols[colName] {
			keepCols = append(keepCols, colName)
		}
	}

	// Drop the specified columns
	return df.Select(keepCols), nil
}
