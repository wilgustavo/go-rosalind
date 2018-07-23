package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Falta indicar archivo")
	}
	stream, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	bases := contarBases(string(stream))
	imprimirResultado(bases)
}

func contarBases(cadena string) map[string]int {
	bases := map[string]int{"A": 0, "C": 0, "G": 0, "T": 0}
	lista := strings.Split(cadena, "")
	for _, item := range lista {
		total, exists := bases[item]
		if exists {
			bases[item] = total + 1
		}
	}
	return bases
}

func imprimirResultado(resultado map[string]int) {
	fmt.Printf("%d %d %d %d", resultado["A"], resultado["C"], resultado["G"], resultado["T"])
}
