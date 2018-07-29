package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounting(t *testing.T) {

	t.Run("Cadena vacía", func(t *testing.T) {
		result := Counting("")
		assert.Equal(t, result, "0 0 0 0")
	})

	t.Run("Un nucleotido de cada uno", func(t *testing.T) {
		result := Counting("ACGT")
		assert.Equal(t, result, "1 1 1 1")
	})

	t.Run("Ejemplo", func(t *testing.T) {
		dna := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
		result := Counting(dna)
		assert.Equal(t, result, "20 12 17 21")
	})

}
