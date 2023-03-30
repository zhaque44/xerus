package decisionanalysis

import (
	"fmt"
	"testing"
	"time"
)

func TestClassifyChurn(t *testing.T) {
	dataset := Dataset{
		Customers: []Customer{
			{
				ID:    1,
				Name:  "John Doe",
				Email: "johndoe@gmail.com",
				Churn: true,
				OrderHistory: []Order{
					{
						ID:          1,
						CustomerID:  1,
						OrderDate:   time.Now().AddDate(0, -9, 0),
						TotalAmount: 50.0,
					},
					{
						ID:          2,
						CustomerID:  1,
						OrderDate:   time.Now().AddDate(0, -10, 0),
						TotalAmount: 75.0,
					},
				},
			},
			{
				ID:    2,
				Name:  "Jane Smith",
				Email: "janesmith@hotmail.com",
				Churn: false,
				OrderHistory: []Order{
					{
						ID:          1,
						CustomerID:  2,
						OrderDate:   time.Now().AddDate(0, -1, 0),
						TotalAmount: 125.0,
					},
					{
						ID:          2,
						CustomerID:  2,
						OrderDate:   time.Now().AddDate(0, -3, 0),
						TotalAmount: 500.0,
					},
				},
			},
			{
				ID:    3,
				Name:  "Bob Jones",
				Email: "bobjones@gmail.com",
				Churn: true,
				OrderHistory: []Order{
					{
						ID:          1,
						CustomerID:  3,
						OrderDate:   time.Now().AddDate(0, -7, 0),
						TotalAmount: 125.0,
					},
				},
			},
		},
	}

	for _, customer := range dataset.Customers {
		if PredictChurn(customer, 200.0, 12) {
			fmt.Printf("%s is likely to churn\n", customer.Name)
		} else {
			fmt.Printf("%s is not likely to churn\n", customer.Name)
		}
	}
}
