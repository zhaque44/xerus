package datapreparation

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/pkg/errors"
)

func LeftJoin(left dataframe.DataFrame, right dataframe.DataFrame, onColumn string) (dataframe.DataFrame, error) {
	if !left.HasCol(onColumn) || !right.HasCol(onColumn) {
		return dataframe.DataFrame{}, errors.Errorf("specified onColumn '%s' not found in both DataFrames", onColumn)
	}

	// Perform cross join to get all combinations
	crossJoined := left.CrossJoin(right)

	// Filter rows to mimic left join behavior
	filteredRows := make([]int, 0)
	leftIndex := left.ColIndex(onColumn)
	rightIndex := right.ColIndex(onColumn)

	for i := 0; i < crossJoined.Nrow(); i++ {
		leftValue := crossJoined.Elem(leftIndex, i)
		rightValue := crossJoined.Elem(rightIndex + left.Ncol(), i)

		if leftValue == rightValue {
			filteredRows = append(filteredRows, i)
		}
	}

	filteredDF := crossJoined.Subset(filteredRows)

	return filteredDF, nil
}
