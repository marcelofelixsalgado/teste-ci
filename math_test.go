package main

import "testing"

func TestSoma(t *testing.T) {

	total := Sum(10, 15)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
