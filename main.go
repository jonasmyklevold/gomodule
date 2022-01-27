package quote

import "fmt"

import (
	"rsc.io/quote"
)

func myquote() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())

}
