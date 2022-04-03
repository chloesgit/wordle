package lib

import "testing"

func TestIsword(t *testing.T) {
	res, err := Isword("penis", "ertyu")
	expected := []string{"Mal placé", "Non existant", "Non existant", "Non existant", "Non existant"}

	if err != nil {
		t.Fail()
	}
	for i, recu := range res {
		if string(recu) != expected[i] {
			t.Error("ca marche po")
		}
	}

	_, err = Isword("penis", "ertyzdqzqzu")
	if err == nil {
		t.Fail()
	}

	res, err = Isword("penis", "eetyu")
	expected = []string{"Mal placé", "Bien placé", "Non existant", "Non existant", "Non existant"}

	if err != nil {
		t.Fail()
	}
	for i, recu := range res {
		if string(recu) != expected[i] {
			t.Error("ca marche po")
		}
	}
	
}
