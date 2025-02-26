package auth

import "testing"

func TestGenerateToken(t *testing.T) {
	tokenTest, err := GenerateJWT("123", "test", "user")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if tokenTest == "" {
		t.Fatalf("Generated token is empty")
	}
}

func TestValidateToken(t *testing.T) {
	tokenTest, _ := GenerateJWT("123", "test", "user")
	ValidationErr := ValidateToken(tokenTest)

	if ValidationErr != nil {
		t.Errorf("ValidateToken() error = %v", ValidationErr)
	}
}
