package us_phone_generator

import (
	"strings"
	"testing"
)

func TestNYPhoneStateGenerate(t *testing.T) {
	t.Run("TestNYPhoneStateGenerate", func(t *testing.T) {
		nyPhone := GenerateByState("NY")
		nyPhone = strings.TrimLeft(nyPhone, "+")

		if len(nyPhone) < 11 {
			t.Fatal("phone number incorrect, want 10 len")
		}
	})
}
