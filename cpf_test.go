package validation

import "testing"

func TestValidationCpf(t *testing.T) {
	type data struct {
		expect bool
		cpf    string
	}
	table := []data{
		{cpf: "58138520018", expect: false},
		{cpf: "21855406020", expect: true},
		{cpf: "95456289016", expect: true},
		{cpf: "74060388017", expect: true},
		{cpf: "12365444444", expect: false},
	}
	for _, item := range table {
		result := Cpf(item.cpf)
		if result != item.expect {
			t.Errorf("expected: %t, result: %t", item.expect, result)
		}
	}
}
