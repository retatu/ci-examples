package main

import "testing"

func TestSome(t *testing.T) {
	total := Soma(15, 10)

	if total != 25 {
		t.Errorf("Resultado Soma inválido. Result: %d Expected: %d", total, 25)
	}

}
