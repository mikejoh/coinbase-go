# coinbase-go

_Golang package to interact with the Coinbase v2 API._

## Example usage

```go
package main

import (
	v2 "coinbase-go/v2"
	"context"
	"fmt"
	"log"
)

func main() {
	config := v2.NewConfig(
		v2.ApiKey("key"),
		v2.Secret("secret"),
	)
	client := v2.NewClient(config)

	exchangeRates, err := client.ExchangeRates(context.Background(), "BTC")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(exchangeRates)
}
```

## TODO
* Rate limit
* Pagination
* Header signing (auth)
* Implement the rest of the API calls

## Resources

[v2 API ref](https://developers.coinbase.com/api/v2#introduction)