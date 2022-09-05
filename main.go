package main

import (
	"fmt"
	"unittests/addresses"
)

func main() {
	addressType := addresses.AddressType("Rua Luiz Oscar de Carvalho")

	fmt.Println(addressType)
}
