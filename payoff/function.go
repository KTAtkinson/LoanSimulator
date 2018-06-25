package payoff

import (
    "github.com/KTAtkinson/LoanSimulator/loan"
)

var (
    MinimumPaymentFunction = NewStrategy("Minimum payment", payMinimumPayments)
)

type StrategyFn func(float64, []*loan.Loan) ([]loan.LoanPayment, float64, error)

type StrategyFunction struct {
    fn StrategyFn
    Name string
}

func NewStrategy(name string, fn StrategyFn) *StrategyFunction {
    return &StrategyFunction{
        Name: name,
        fn: fn,
    }
}

func (f *StrategyFunction) Call(purseValue float64, loans []*loan.Loan) ([]loan.LoanPayment, float64, error) {
    return f.fn(purseValue, loans)
}
