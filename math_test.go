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
func TestSub(t *testing.T) {
	total := sub(5, 5)

	if total != 0 {
		t.Errorf("Resultado Soma inválido. Result: %d Expected: %d", total, 25)
	}

}
func Test1(t *testing.T) {
	total := teste1(5, 5)

	if total != 0 {
		t.Errorf("Resultado Soma inválido. Result: %d Expected: %d", total, 25)
	}

}
func Test2(t *testing.T) {
	total := teste2(5, 5)

	if total != 0 {
		t.Errorf("Resultado Soma inválido. Result: %d Expected: %d", total, 25)
	}

}
