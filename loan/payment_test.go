package loan

import (
    "testing"
)

func TestPaymentAmount(t *testing.T) {
    newPaymentSpec := []struct {
        interest float64
        principal float64
        amount float64
    } {
        {1, 2, 3},
        {0, 1, 1},
        {1, 0, 1},
    }

    for _, spec := range(newPaymentSpec) {
        newPayment := NewPayment(spec.interest, spec.principal)
        amount := newPayment.Amount()
        if amount != spec.amount {
            t.Errorf("Expected amount paid to be %v, found %v.", spec.amount, amount)
        }
    }
}
