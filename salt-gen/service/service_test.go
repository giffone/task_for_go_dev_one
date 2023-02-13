package service

import "testing"

func TestGenerate(t *testing.T) {
	salt, err := Generate(12)
	if err != nil {
		t.Fatalf("salt generate: %s", err)
	}
	if len(salt) != 12 {
		t.Fatalf("length salt is: %d", len(salt))
	}
}
