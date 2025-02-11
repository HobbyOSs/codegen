package ocode

import (
	"strings"
	"errors"
)

// CodegenClient インターフェースの定義
type CodegenClient interface {
	Emit(line string) error
	EmitAll(text string) error
	Exec() (float64, error)
}

// ocodeClient 構造体の定義
type ocodeClient struct {
	// 必要なフィールドを定義
}

// NewCodegenClient は新しい CodegenClient を返す
func NewCodegenClient() CodegenClient {
	return &ocodeClient{
		// フィールドの初期化
	}
}

// Emit メソッドの実装
func (c *ocodeClient) Emit(line string) error {
	// 単一行の命令を処理するロジックを実装
	return nil
}

// EmitAll メソッドの実装
func (c *ocodeClient) EmitAll(text string) error {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	for _, line := range lines {
		if err := c.Emit(strings.TrimSpace(line)); err != nil {
			return err
		}
	}
	return nil
}

// Exec メソッドの実装
func (c *ocodeClient) Exec() (float64, error) {
	// 実行ロジックを実装
	return 0, errors.New("not implemented")
}
