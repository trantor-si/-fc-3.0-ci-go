package main

import "testing"

func TestSum(t *testing.T) {
	esperado := 30
	total := sum(15, 15)
	if total != esperado {
		t.Errorf("Resultado da soma de 15+15 é inválido. Resultado %d. Esperado %d", total, esperado)
	}
}

func TestSub(t *testing.T) {
	esperado := 15
	total := sub(30, 15)
	if total != esperado {
		t.Errorf("Resultado da subtração de 30-15 é inválido. Resultado %d. Esperado %d", total, esperado)
	}
}

func TestMult(t *testing.T) {
	esperado := 20
	total := mult(10, 2)
	if total != esperado {
		t.Errorf("Resultado da multiplicação de 10*2 é inválido. Resultado %d. Esperado %d", total, esperado)
	}
}
func TestDiv(t *testing.T) {
	esperado := 10
	total := div(20, 2)
	if total != esperado {
		t.Errorf("Resultado da divisão de 20/2 é inválido. Resultado %d. Esperado %d", total, esperado)
	}
}
