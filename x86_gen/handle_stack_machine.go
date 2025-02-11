package x86_gen

import (
	"fmt"

	"github.com/HobbyOSs/codegen/variantstack"
)

func handleL(args []string, vs *variantstack.VariantStack) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("L requires 1 argument")
	}
	vs.Push(args[0])
	return nil, nil
}
