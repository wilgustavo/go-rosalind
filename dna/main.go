package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func formatResultado(datos map[string]int) string {
	return fmt.Sprintf("%d %d %d %d", datos["A"], datos["C"], datos["G"], datos["T"])
}

func contar(dna string) (result map[string]int) {
	result = map[string]int{}
	for _, base := range dna {
		result[string(base)]++
	}
	return
}

// Counting return four integers (separated by spaces) counting the respective number of
// times that the symbols 'A', 'C', 'G', and 'T' occur in s.
func Counting(dna string) string {
	return formatResultado(contar(dna))
}

func leerDNA(archivo string) {
	datos, err := ioutil.ReadFile(archivo)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(Counting(string(datos)))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Debe indicar el nombre del archivo")
		os.Exit(0)
	}
	leerDNA(os.Args[1])
}
