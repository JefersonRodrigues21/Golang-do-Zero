package main

import (
	"fmt"
	steph "strings"
)

func main() {
	fmt.Println("Hello, world!")

	// sem alterar o nome do pacote,
	// usaríamos chamando strings.Func()
	// strings.Slipt()
	fmt.Println(steph.Split("steph", ""))
}
// TIPOS:
// boll (true/false)
// string - sequência de bytes
// int - numeros inteiros
// float (float64/float32) - decimal