package client

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

func TestExec(t *testing.T) {
	client := NewCodegenClient()

	err := client.EmitAll("LN 5\nLN EAX\nADD")
	if err != nil {
		t.Fatalf("EmitAll failed: %v", err)
	}

	_, err = client.Exec()
	if err != nil {
		t.Errorf("Exec failed: %v", err)
	}

	// TODO: ここでスタックの内部状態を確認するテストを追加
	// 例えば、スタックが期待通りの状態になっているかを確認
}

func TestEmitAll(t *testing.T) {
	client := NewCodegenClient()

	err := client.EmitAll("LN 5\nLN EAX\nADD")
	if err != nil {
		t.Errorf("EmitAll failed: %v", err)
	}
}
