package main

import (
	"flag"
	"fmt"
)

func main() {
	var nflag = flag.Int("n", 1234, "Used to define a number")
	flag.Parse()
	fmt.Println(*nflag)
}
