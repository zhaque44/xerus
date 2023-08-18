package datapreparation

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/pkg/errors"
)

func LeftJoin(left dataframe.DataFrame, right dataframe.DataFrame, onColumn string) (dataframe.DataFrame, error) {
	// Check if the join column exists in both DataFrames
	if !left.HasCol(onColumn) || !right.HasCol(onColumn) {
		return dataframe.DataFrame{}, errors.Errorf("specified onColumn '%s' not found in both DataFrames", onColumn)
	}

	leftIndex := left.ColIndex(onColumn)
	rightIndex := right.ColIndex(onColumn)

	// Check if the join columns have the same data type
	if left.ColTypes()[leftIndex] != right.ColTypes()[rightIndex] {
		return dataframe.DataFrame{}, errors.New("join columns have different data types")
	}

	// Perform cross join to create the initial cross-joined DataFrame
	crossJoined := left.CrossJoin(right)

	// Initialize a slice to store indices of rows that match the join condition
	matchingRows := make([]int, 0)

	// Filter rows based on matching join condition
	for i := 0; i < crossJoined.Nrow(); i++ {
		leftValue := crossJoined.Elem(leftIndex, i)
		rightValue := crossJoined.Elem(rightIndex + left.Ncol(), i)

		if leftValue == rightValue {
			matchingRows = append(matchingRows, i)
		}
	}

	// Subset the DataFrame with the matching rows
	filteredDF := crossJoined.Subset(matchingRows)

	return filteredDF, nil
}
