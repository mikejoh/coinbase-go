# coinbase-go

![Test workflow](https://github.com/mikejoh/coinbase-go/actions/workflows/test.yml/badge.svg)

An alternative Golang package to interact with the [Coinbase](https://www.coinbase.com/) **v2 API**.

_Note that this project is still work in progress!_

## Installation
Install:
```
go get github.com/mikejoh/coinbase-go
```
Import (with alias):
```go
import cb "github.com/mikejoh/coinbase-go"
```
## Examples

Instantiate client with config (using import alias):
```go
package main

import cb "github.com/mikejoh/coinbase-go"

func main() {
	config := cb.NewConfig(
		cb.ApiKey("key"),
		cb.Secret("secret"),
	)

	client := cb.NewClient(config)
}
```

### Exchange rates
```go
rates, err := client.ExchangeRates(context.Background(), "BTC")
if err != nil {
	log.Fatal(err)
}

fmt.Println(rates)
```
### Currencies
```go
currencies, err := client.Currencies(context.Background())
if err != nil {
	log.Fatal(err)
}

fmt.Println(currencies)
```

### Prices
```go
prices, err := client.Prices(context.Background(), "BTC-SEK", "sell")
if err != nil {
	log.Fatal(err)
}

fmt.Println(prices)
```

### Time
```go
time, err := client.Time(context.Background())
if err != nil {
	log.Fatal(err)
}

fmt.Println(time)
```

## Coinbase CLI

1. Build:
```
make client
```
The binary is created in `./bin`.

2. Example usage:
```
./bin/cb currencies
./bin/cb prices --pair BTC-SEK --type buy
./bin/cb exchange-rates
./bin/cb exchange-rates --currency SEK
./bin/cb time
```
## Resources

[v2 API ref](https://developers.coinbase.com/api/v2#introduction)
## Todo
* [ ] Add code documentation
* [ ] Rate limit
* [ ] Pagination
* [ ] Header signing (auth)
* [ ] Implement Wallet endpoints
* [ ] Implement a more generic way of making client API calls, DRY.

