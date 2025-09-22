package syntax

import "fmt"

// Interface defines the behavior
type PaymentMethod interface {
	Pay(amount float64)
}

// Credit Card struct
type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card [%s]\n", amount, c.CardNumber)
}

// PayPal struct
type PayPal struct {
	Email string
}

func (p PayPal) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal account [%s]\n", amount, p.Email)
}

// Bitcoin struct
type Bitcoin struct {
	WalletAddress string
}

func (b Bitcoin) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Bitcoin wallet [%s]\n", amount, b.WalletAddress)
}

func OopAction() {
	// Different payment methods
	methods := []PaymentMethod{
		CreditCard{CardNumber: "1234-5678-9999"},
		PayPal{Email: "user@example.com"},
		Bitcoin{WalletAddress: "1A1zP1..."},
	}

	// Process payments polymorphically
	for _, m := range methods {
		m.Pay(100.0) // each behaves differently
	}
}
