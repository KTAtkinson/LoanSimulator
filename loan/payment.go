package loan

type LoanPayment struct {
	interest  float64
	principle float64
}

func NewPayment(interestToPay, principleToPay float64) LoanPayment {
    return LoanPayment{
        interest: interestToPay,
        principle: principleToPay,
    }
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
