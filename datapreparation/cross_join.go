


func LeftJoin(left DataFrame, right DataFrame, onColumn string) (DataFrame, error) {
    // Check if the onColumn exists in both DataFrames
    if !left.HasCol(onColumn) || !right.HasCol(onColumn) {
        return DataFrame{}, fmt.Errorf("specified onColumn '%s' not found in both DataFrames", onColumn)
    }
    
    // Find the indices of the onColumn in both DataFrames
    leftIndex := left.ColIndex(onColumn)
    rightIndex := right.ColIndex(onColumn)

    // Initialize the resulting DataFrame
    var resultCols []series.Series
    for _, col := range left.Columns() {
        resultCols = append(resultCols, col.Empty())
    }
    for _, col := range right.Columns() {
        if col.Name() != onColumn {
            resultCols = append(resultCols, col.Empty())
        }
    }

    // Iterate through the left DataFrame and perform the left join
    for i := 0; i < left.Nrow(); i++ {
        leftValue := left.Elem(leftIndex, i)
        found := false

        // Search for matching rows in the right DataFrame
        for j := 0; j < right.Nrow(); j++ {
            rightValue := right.Elem(rightIndex, j)
            if leftValue == rightValue {
                found = true

                // Append values from the left DataFrame
                for _, col := range left.Columns() {
                    resultCols[col.Index()].Append(col.Elem(i))
                }

                // Append values from the right DataFrame (excluding the onColumn)
                for _, col := range right.Columns() {
                    if col.Name() != onColumn {
                        resultCols[col.Index()+left.Ncol()].Append(col.Elem(j))
                    }
                }
            }
        }

        // If no match found, append null values for right DataFrame columns
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

    return New(resultCols...), nil
}
