package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("Hash returned empty.")
	}

	if hash == "password" {
		t.Error("Hash is same as password.")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !ComparePasswords(hash, []byte("password")) {
		t.Error("Expected password to match hash.")
	}
	if ComparePasswords(hash, []byte("notthepassword")) {
		t.Error("Expected passwords to not match.")
	}
}
