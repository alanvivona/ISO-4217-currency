# ISO-4217-currency

Manage currency in ISO 4217 format.

### Download:  
`go get -u github.com/alanvivona/ISO-4217-currency`

### Example:  
```go
package main

import (
	"fmt"

	isocur "github.com/alanvivona/ISO-4217-currency"
)

func main() {

	var c isocur.Currency

	// Get currency by code
	c, err := isocur.GetByCode("GBP") // OK
	if err != nil {
		panic(":(")
	}
	fmt.Printf("British Pound: %s\n", c)

	c, err = isocur.GetByCode("123") // ERR: Currency code not found
	if err != nil {
		fmt.Printf("Unknown currecny symbol\n")
	}

	// Get currency by name
	hkd := isocur.HongKongDollar
	y := isocur.JapanYen

	ausd := isocur.AustraliaDollar
	usd, _ := isocur.GetByCode("USD")

	fmt.Printf("Hong Kong Dollar: %s\n", hkd)
	fmt.Printf("Japan Yen: %s\n", y)
	fmt.Printf("Australian Dollar: %s\n", ausd)
	fmt.Printf("United States Dollar: %s\n", usd)
}
```

    > go run main.go
    British Pound: GBP
    Unknown currecny symbol
    Hong Kong Dollar: HKD
    Japan Yen: JPY
    Australian Dollar: AUD
    United States Dollar: USD
