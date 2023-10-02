package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprartv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprartv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)
}