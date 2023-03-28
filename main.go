package main

import (
	"fmt"
	"os"
	"xerus/datapreparation"

	"github.com/pkg/errors"
)

func main() {
	file, err := os.Open("datapreparation/data.csv")
	if err != nil {
		panic(errors.Wrap(err, "failed to open CSV file"))
	}
	defer file.Close()

	// Load CSV into a DataFrame
	df, err := datapreparation.LoadCSV(file)
	if err != nil {
		panic(errors.Wrap(err, "failed to load CSV into DataFrame"))
	}

	// Print the first 5 rows of the DataFrame
	fmt.Println(df)
}
