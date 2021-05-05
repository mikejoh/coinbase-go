# coinbase-go

_A Golang package to interact with the [Coinbase](https://www.coinbase.com/) **v2 API**._

_Note that this project is still work in progress!_
## Examples

Instantiate client with config:
```go
func main() {
	config := v2.NewConfig(
		v2.ApiKey("key"),
		v2.Secret("secret"),
	)

	client := v2.NewClient(config)
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
## Resources

[v2 API ref](https://developers.coinbase.com/api/v2#introduction)
## TODO
* [ ] Rate limit
* [ ] Pagination
* [ ] Header signing (auth)
* [ ] Implement Wallet endpoints
* [ ] Implement a more generic way of making client API calls, DRY.

