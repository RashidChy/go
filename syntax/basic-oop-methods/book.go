package methods

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) print() {
	// b is a copy of the original `book` value here.
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}
