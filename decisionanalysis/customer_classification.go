package decisionanalysis

import "time"

type Customer struct {
	ID           int
	Name         string
	Email        string
	Churn        bool
	OrderHistory []Order
}

type Order struct {
	ID          int
	CustomerID  int
	OrderDate   time.Time
	TotalAmount float64
}

type Dataset struct {
	Customers []Customer
}

func PredictChurn(customer Customer, churnThresholdAmount float64, churnThresholdMonths int) bool {
	var totalAmount float64
	for _, order := range customer.OrderHistory {
		totalAmount += order.TotalAmount
	}

	if totalAmount < churnThresholdAmount || time.Since(customer.OrderHistory[len(customer.OrderHistory)-1].OrderDate).Hours()/24/30 > float64(churnThresholdMonths) {
		return true
	} else {
		return false
	}
}
