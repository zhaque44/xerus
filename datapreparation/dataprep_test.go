package datapreparation

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/stretchr/testify/assert"
)

func TestLoadCSV(t *testing.T) {
	file, err := os.Open("data.csv")
	if err != nil {
		t.Fatalf("Failed to open CSV file: %v", err)
	}

	defer file.Close()

	df, err := LoadCSV(file)

	if err != nil {
		t.Fatalf("Failed to load CSV into DataFrame: %v", err)
	}

	if df.Nrow() == 0 {
		t.Fatalf("DataFrame is empty")
	}

	if df.Ncol() == 0 {
		t.Fatalf("DataFrame has no columns")
	}
}

func TestDropCols(t *testing.T) {
	file, err := os.Open("data.csv")
	df := dataframe.ReadCSV(file, dataframe.DetectTypes(true))

	newDf, err := DropCols(df, []string{"age"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newDf)

	assert.Equal(t, newDf.Ncol(), 3)
	// assert age column is not present
	assert.NotContains(t, newDf.Names(), "age")
}
