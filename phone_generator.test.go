package us_phone_generator

import (
	"strings"
	"testing"
)

func TestNYPhoneStateGenerate(t *testing.T) {
	t.Run("TestNYPhoneStateGenerate", func(t *testing.T) {
		phoneNumber, err := GenerateByState("NY")
		if err != nil {
			t.Fatalf("phone number incorrect, err: %v", err)
		}

		phoneNumber = strings.TrimLeft(phoneNumber, "+")

		if len(phoneNumber) < 11 {
			t.Fatal("phone number incorrect, want 10 len")
		}
	})

	t.Run("TestNewMassachusettsPhoneStateGenerate", func(t *testing.T) {
		phoneNumber, err := GenerateByState("massachusetts")
		if err != nil {
			t.Fatalf("phone number incorrect, err: %v", err)
		}

		phoneNumber = strings.TrimLeft(phoneNumber, "+")

		if len(phoneNumber) < 11 {
			t.Fatal("phone number incorrect, want 10 len")
		}
	})
}
