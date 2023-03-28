package datapreparation

import (
	"os"
	"testing"
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
