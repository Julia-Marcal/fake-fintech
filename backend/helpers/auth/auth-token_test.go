package auth

import "testing"

func TestGenerateToken(t *testing.T) {
	tokenTest, err := GenerateJWT("test@gmail.com", "test")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if tokenTest == "" {
		t.Fatalf("Generated token is empty")
	}
}

func TestValidateToken(t *testing.T) {
	tokenTest, _ := GenerateJWT("test@gmail.com", "test")
	ValidationErr := ValidateToken(tokenTest)

	if ValidationErr != nil {
		t.Errorf("ValidateToken() error = %v", ValidationErr)
	}
}
