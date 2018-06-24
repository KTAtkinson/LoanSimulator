package loan

import (
	"errors"
	"math"
)

type Loan struct {
	principle            float64
	amount               float64
	monthlyPayment       float64
	term                 uint
	annualPercentageRate float64
}

func NewLoan(amount float64, interestRate float64, term uint) (Loan, error) {
	if amount <= 0 {
		return Loan{}, errors.New("Starting principle must be greater than 0.")
	}
	if interestRate <= 0 {
		return Loan{}, errors.New("Interest rate must be greater than 0.")
	}

	interestPerPayment := interestRate / 12
	discountFactor := (math.Pow(1+interestPerPayment, float64(term)) - 1) / (interestPerPayment * math.Pow(1+interestPerPayment, float64(term)))
	monthlyPayment := float64(amount) / discountFactor

	loan := Loan{
		principle:            amount,
		amount:               amount,
		monthlyPayment:       monthlyPayment,
		term:                 term,
		annualPercentageRate: interestRate,
	}
	return loan, nil
}

func (l *Loan) Amount() float64 {
	return l.amount
}

func (l *Loan) Term() uint {
	return l.term
}

func (l *Loan) AnnualPrecentageRate() float64 {
	return l.annualPercentageRate
}

func (l *Loan) CalculateNextPayment() (float64, float64, error) {
	if l.principle <= 0 {
		return 0, 0, errors.New("Loan paid in full")
	}

	interestDue := l.principle * (l.annualPercentageRate / 12)
	principleDue := l.monthlyPayment - interestDue
	return interestDue, principleDue, nil
}

func (l *Loan) ApplyMontlyPayment(payment LoanPayment) (float64, error) {
	if payment.Amount() < l.monthlyPayment {
		return payment.Amount(), NewPaymentError(payment.Amount(), l.principle, InsufficiantFunds)
	}

	return l.ApplyPayment(payment), nil
}

func (l *Loan) ApplyPayment(payment LoanPayment) float64 {
	paymentAmount := payment.Amount()
	if paymentAmount > l.principle {
		l.principle = 0
		return paymentAmount - l.principle
	}

	l.principle -= payment.Principle()
	return 0
}
