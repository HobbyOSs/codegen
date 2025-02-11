package x86_gen

import (
	"testing"

	"github.com/HobbyOSs/codegen/ocode"
)

func TestGenerateX86(t *testing.T) {
	tests := []struct {
		name     string
		ocodes   []ocode.Ocode
		expected []byte
	}{
		{
			name: "DB",
			ocodes: []ocode.Ocode{
				{Kind: ocode.OpDB, Operands: []interface{}{0x02, 0xe0}},
			},
			expected: []byte{0x02, 0xe0},
		},
		{
			name: "DW",
			ocodes: []ocode.Ocode{
				{Kind: ocode.OpDW, Operands: []interface{}{0x1234}},
			},
			expected: []byte{0x34, 0x12},
		},
		{
			name: "DD",
			ocodes: []ocode.Ocode{
				{Kind: ocode.OpDD, Operands: []interface{}{0x12345678}},
			},
			expected: []byte{0x78, 0x56, 0x34, 0x12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateX86(tt.ocodes)
			if !equal(result, tt.expected) {
				t.Errorf("got %v, expected %v", result, tt.expected)
			}
		})
	}
}
