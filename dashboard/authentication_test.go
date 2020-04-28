package dashboard

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken()

	if len(token) != 60 {
		t.Errorf("GenerateToken() failed.")
	}
}

func TestHashToken(t *testing.T) {
	token := "hello"
	authenticationHash = HashToken(token)

	if !ChecksToken(token) {
		t.Errorf("HashToken() failed.")
	}
}

func TestSaveHash(t *testing.T) {
	fileName = "../" + fileName
	SaveHash("hey")

	if !AuthenticationFileExists() {
		t.Errorf("SaveHash() failed.")
	}

	Authenticate()

	if string(authenticationHash) != "hey" {
		t.Errorf("Authenticate() failed.")
	}
}
