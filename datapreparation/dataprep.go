package datapreparation

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func LoadCSV(file *os.File) (*dataframe.DataFrame, error) {
	r := bufio.NewReader(file)

	df := dataframe.ReadCSV(r)

	if err := df.Err; err != nil {
		return nil, fmt.Errorf("error reading CSV: %s", err.Error())
	}

	if len(df.Names()) == 0 {
		return nil, fmt.Errorf("CSV file has no header row")
	}

	return &df, nil
}
