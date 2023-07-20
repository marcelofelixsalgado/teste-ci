package main

import "testing"

func TestSoma(t *testing.T) {

	total := Sum(1, 2)

	if total != 3 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
