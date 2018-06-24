package loan

import (
	"errors"
	"fmt"
)

var (
	InsufficiantFunds = errors.New("Insufficiant funds to make payment.")
)

type PaymentError struct {
	Err                error
	PaymentAmount      float64
	OutstandingBalance float64
}

func NewPaymentError(paymentAmount float64, balance float64, err error) *PaymentError {
	return &PaymentError{
		Err:                err,
		PaymentAmount:      paymentAmount,
		OutstandingBalance: balance,
	}
}

func (e *PaymentError) Error() string {
	return fmt.Sprintf("Unable to apply %g to balance of %d. %s",
		e.PaymentAmount,
		e.OutstandingBalance,
		e.Err.Error())
}
