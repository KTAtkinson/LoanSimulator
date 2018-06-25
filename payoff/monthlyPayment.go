package payoff

import (
    "github.com/KTAtkinson/LoanSimulator/loan"
)

func payMinimumPayments(purseValue float64, loans []*loan.Loan) ([]loan.LoanPayment, float64, error) {
    var payments []loan.LoanPayment
    for _, cur := range loans {
        interestDue, principalDue, err := cur.CalculateNextPayment()
        if err != nil {
            return payments, purseValue, err
        }

        payment := loan.NewPayment(interestDue, principalDue)
        purseValue, err = cur.ApplyMonthlyPayment(payment)
        if err != nil {
            return payments, purseValue, nil
        }

        payments = append(payments, payment)
    }

    return payments, purseValue, nil
}
