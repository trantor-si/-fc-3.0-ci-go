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
		t.Errorf("Resultado da soma de 30-15 é inválido. Resultado %d. Esperado %d", total, esperado)
	}
}
