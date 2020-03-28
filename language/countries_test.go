package language

import (
	"testing"
)

func TestSerializeCountries(t *testing.T) {
	countries := SerializeCountries()
	excepted := "Aruba"

	if countries[0].CommonName != excepted {
		t.Errorf("SerializeCountries() failed, excepted %s got %s.", excepted, countries[0].CommonName)
	}
}

func TestFindCountry(t *testing.T) {
	sentence := "What is the capital of France please"
	excepted := "France"
	country := FindCountry(sentence).CommonName

	if excepted != country {
		t.Errorf("FindCountry() failed, excepted %s got %s.", excepted, country)
	}
}
