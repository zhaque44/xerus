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

	// Check if the join columns have the same data type
	if left.ColTypes()[leftIndex] != right.ColTypes()[rightIndex] {
		return dataframe.DataFrame{}, errors.New("join columns have different data types")
	}

	var resultCols []dataframe.Series
	for _, col := range left.Columns() {
		resultCols = append(resultCols, col.Empty())
	}
	for _, col := range right.Columns() {
		if col.Name() != onColumn {
			resultCols = append(resultCols, col.Empty())
		}
	}

	for i := 0; i < left.Nrow(); i++ {
		leftValue := left.Elem(leftIndex, i)
		found := false

		for j := 0; j < right.Nrow(); j++ {
			rightValue := right.Elem(rightIndex, j)
			if leftValue == rightValue {
				found = true

				for _, col := range left.Columns() {
					resultCols[col.Index()].Append(col.Elem(i))
				}

				for _, col := range right.Columns() {
					if col.Name() != onColumn {
						resultCols[col.Index()+left.Ncol()].Append(col.Elem(j))
					}
				}
			}
		}

		if !found {
			for _, col := range left.Columns() {
				resultCols[col.Index()].Append(col.Elem(i))
			}
			for _, col := range right.Columns() {
				if col.Name() != onColumn {
					resultCols[col.Index()+left.Ncol()].Append(nil)
				}
			}
		}
	}

	return dataframe.New(resultCols...), nil
}
	
