package client

import (
	"strings"

	"github.com/HobbyOSs/codegen/ocode"
	"github.com/HobbyOSs/codegen/x86_gen"
)

// CodegenClient インターフェースの定義
type CodegenClient interface {
	Emit(line string) error
	EmitAll(text string) error
	Exec() (float64, error)
}

// ocodeClient 構造体の定義
type ocodeClient struct {
	ocodes []ocode.Ocode
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
	ocode, err := parseLineToOcode(line)
	if err != nil {
		return err
	}
	c.ocodes = append(c.ocodes, ocode)
	return nil
}

// TODO parseLineToOcode 実装
func parseLineToOcode(line string) (ocode.Ocode, error) {
	panic("unimplemented")
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
	_ = x86_gen.GenerateX86(c.ocodes)
	// TODO: ここで machineCode を使って何らかの処理を行う
	return 0, nil // 適切な戻り値に変更
}
