# ISO-4217-currency

Manage currency in ISO 4217 format.

```go
    var c currency.Currency
    var err error

    // Get currency by code
    c, err = GetByCode("GBP")  // OK
    c, err = GetByCode("123")  // ERR: Currency code not found

    // Get currency by name
    c = currency.HongKongDollar
    c = currency.JapanYen
```