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
