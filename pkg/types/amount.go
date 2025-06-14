package types

import (
	"fmt"
	"strconv"
)

// Amount represents a monetary amount with currency information
// Equivalent to PHP Amount class
type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

// NewAmount creates a new Amount instance with proper formatting
func NewAmount(amount float64, currency string) *Amount {
	if currency == "" {
		currency = "IDR"
	}

	// Format amount to 2 decimal places as string
	value := fmt.Sprintf("%.2f", amount)

	return &Amount{
		Value:    value,
		Currency: currency,
	}
}

// Float64 returns the amount as a float64 value
func (a *Amount) Float64() (float64, error) {
	return strconv.ParseFloat(a.Value, 64)
}

// String returns a formatted string representation of the amount
func (a *Amount) String() string {
	return fmt.Sprintf("%s %s", a.Value, a.Currency)
}
