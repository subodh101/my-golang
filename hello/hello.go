package hello

import (
	"rsc.io/quote"
	QuoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quote.Hello()
}

func Proverb() string {
	return QuoteV3.Concurrency()
}
