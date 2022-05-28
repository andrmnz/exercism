package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	const rate1, rate2, rate3, rate4 float32 = 3.213, 0.5, 1.621, 2.475

	if balance < 0 {
		return rate1
	} else if balance < 1000 {
		return rate2
	} else if balance < 5000 {
		return rate3
	} else {
		return rate4
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) / 100 * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return float64(Interest(balance)) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int

	for balance < targetBalance {
		years++
		balance = AnnualBalanceUpdate(balance)
	}

	return years
}
