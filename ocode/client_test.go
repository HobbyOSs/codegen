package ocode

import (
	"testing"
)

func TestEmit(t *testing.T) {
	client := NewCodegenClient()

	err := client.Emit("LN 5")
	if err != nil {
		t.Errorf("Emit failed: %v", err)
	}
}

func TestEmitAll(t *testing.T) {
	client := NewCodegenClient()

	err := client.EmitAll("LN 5\nLN EAX\nADD")
	if err != nil {
		t.Errorf("EmitAll failed: %v", err)
	}
}
