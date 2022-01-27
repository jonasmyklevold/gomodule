package gomodule

import "fmt"
import "rsc.io/quote"

func gomodule() {
	fmt.Println(getGlass())
}

func getGlass() string {
	fmt.Println(quote.Glass())
}

func getGo() string {
	return quote.Go()
}

func getHello() string {
	return getHello()
}

func getOpt() string {
	return quote.Go()
}
