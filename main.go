ackage main

import (
	"fmt"
	"github.com/jonasmyklevold/gomodule/state"
)

func main() {
	fmt.Println(state.ViewState())
	state.PutInKylling()
	fmt.Println(state.ViewState())
}