package quote

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println(getGlass())
}

func getGlass() string {
	return quote.Glass()
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
