package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 10)

	if total != 25 {
		t.Errorf("Resultado Soma inválido. Result: %d Expected: %d", total, 25)
	}

}
func TestMulti(t *testing.T) {
	total := multi(5, 5)

	if total != 25 {
		t.Errorf("Resultado Soma inválido. Result: %d Expected: %d", total, 25)
	}

}
