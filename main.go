package main

import (
	"fmt"
	"os"
	"xerus/datapreparation"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	df, err := datapreparation.LoadCSV(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print the first 5 rows of the DataFrame
	fmt.Println(df)
}
