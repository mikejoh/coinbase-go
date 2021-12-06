package coinbase

import "fmt"

type errorResponse struct {
	Errors []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"errors"`
	Code int
	URL  string
}

func (er errorResponse) Error() string {
	return fmt.Sprintf("HTTP %d %s calling %s", er.Code, er.Errors[0].Message, er.URL)
}
