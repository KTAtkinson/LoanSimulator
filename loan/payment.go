package loan

type LoanPayment struct {
	interest  float64
	principle float64
	amount    float64
}

func NewPayment(interestDue, principleDue, paymentAmount float64) (LoanPayment, error) {
	amountDue := interestDue + principleDue
	if amountDue > paymentAmount {
		return LoanPayment{}, NewPaymentError(paymentAmount, amountDue, InsufficiantFunds)
	}

	var principleToPay float64
	if principleDue == 0 {
		principleToPay = paymentAmount
	} else {
		principleToPay = paymentAmount
	}

	payment := LoanPayment{
		interest:  interestDue,
		principle: principleToPay,
	}

	return payment, nil
}

func (p LoanPayment) Interest() float64 {
	return p.interest
}

func (p LoanPayment) Principle() float64 {
	return p.principle
}

func (p LoanPayment) Amount() float64 {
	return p.interest + p.principle
}
