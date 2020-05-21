package language

import (
	"testing"
)

func TestSerializeCountries(t *testing.T) {
	countries := SerializeCountries()
	excepted := "United Kingdom"

	if countries[0].Name["en"] != excepted {
		t.Errorf("SerializeCountries() failed, excepted %s got %s.", excepted, countries[0].Name["en"])
	}
}

func TestFindCountry(t *testing.T) {
	sentence := "What is the capital of France please"
	excepted := "France"
	country := FindCountry("en", sentence).Name["en"]

	if excepted != country {
		t.Errorf("FindCountry() failed, excepted %s got %s.", excepted, country)
	}
}
