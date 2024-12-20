package security

import (
	"bytes"
	"testing"

	env "github.com/Julia-Marcal/reusable-api/config/env"
)

func TestDeriveKey(t *testing.T) {
	var mySalt = []byte("MyS4Lt")
	hashedPass, salt, err := DeriveKey("mypassword", mySalt)

	if err != nil {
		t.Errorf("function returned a error %s", err)
	}

	if salt == nil {
		t.Errorf("Salt is being returned empty")
	}

	if hashedPass == nil {
		t.Errorf("Hashed password is being returned empty")
	}

	if !bytes.Equal(salt, mySalt) {
		t.Errorf("Salt is being returned different than the original")
	}
}

func TestLoginSystem(t *testing.T) {
	salt := env.SetSalt()

	hashedPass, _, _ := DeriveKey("mypassword", salt)
	err := LoginSystem("mypassword", string(hashedPass))

	if err != nil {
		t.Errorf("function returned a error %s", err)
	}

}
