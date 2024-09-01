package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Error("Could not create JWT.")
	}

	if token == "" {
		t.Error("Token is empty.")
	}
}
