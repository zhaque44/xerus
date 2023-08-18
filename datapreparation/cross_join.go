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

	leftIndex := left.ColIndex(onColumn)
	rightIndex := right.ColIndex(onColumn)

	if left.ColTypes()[leftIndex] != right.ColTypes()[rightIndex] {
		return dataframe.DataFrame{}, errors.New("join columns have different data types")
	}

	crossJoined := left.CrossJoin(right)

	matchingRows := make([]int, 0)

	for i := 0; i < crossJoined.Nrow(); i++ {
		leftValue := crossJoined.Elem(leftIndex, i)
		rightValue := crossJoined.Elem(rightIndex + left.Ncol(), i)

		if leftValue == rightValue {
			matchingRows = append(matchingRows, i)
		}
	}

	filteredDF := crossJoined.Subset(matchingRows)
	return filteredDF, nil
}
