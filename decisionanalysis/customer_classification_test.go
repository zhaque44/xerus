package decisionanalysis

import (
	"fmt"
	"testing"
)

func TestClassifyChurn(t *testing.T) {
	dataset := Dataset{
		Customers: []Customer{
			{
				ID:     1,
				Name:   "John Doe",
				Email:  "johndoe@gmail.com",
				Churn:  false,
				Orders: 10,
			},
			{
				ID:     2,
				Name:   "Jane Smith",
				Email:  "janesmith@yahoo.com",
				Churn:  true,
				Orders: 5,
			},
			{
				ID:     3,
				Name:   "Bob Johnson",
				Email:  "bobjohnson@gmail.com",
				Churn:  false,
				Orders: 8,
			},
		},
	}

	for _, customer := range dataset.Customers {
		if ClassifyChurn(customer) {
			fmt.Printf("Customer %d is likely to churn\n", customer.ID)
		} else {
			fmt.Printf("Customer %d is not likely to churn\n", customer.ID)
		}
	}
}
