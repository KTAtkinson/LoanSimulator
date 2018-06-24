package loan

import (
	"math"
	"testing"
)

func TestNewLoan(t *testing.T) {
	loanSpecs := []struct {
		amount               float64
		term                 uint
		annualPercentageRate float64
		monthlyPayment       float64
	}{
		{1230000, 360, .04125, 5961.19},
		{100250, 360, .10, 879.77},
		{6000, 120, .10, 79.29},
		{21000000000000, 1200, .0475, 83857342748.66},
	}

	for _, example := range loanSpecs {
		newLoan, _ := NewLoan(example.amount, example.annualPercentageRate, example.term)
		if math.Abs(newLoan.monthlyPayment-example.monthlyPayment) > (.0001 * example.monthlyPayment) {
			t.Errorf("Expected monthly payment to be %v, is %v.", example.monthlyPayment, newLoan.monthlyPayment)
		}
		if newLoan.amount != example.amount {
			t.Errorf("Expected new loan amount to be %v, is %v.", example.amount, newLoan.amount)
		}
		if newLoan.principle != example.amount {
			t.Errorf("Expected new loan principle to be %v, is %v.", example.amount, newLoan.principle)
		}
		if newLoan.term != example.term {
			t.Errorf("Expected new loan term to be %v, is %v.", example.term, newLoan.term)
		}
		if newLoan.annualPercentageRate != example.annualPercentageRate {
			t.Errorf("Expected percentage rate %v, is %v.", example.annualPercentageRate, newLoan.annualPercentageRate)
		}
	}
}
