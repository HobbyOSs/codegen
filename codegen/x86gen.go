package backend

import "github.com/HobbyOSs/codegen/ocode"

func GenerateX86(ocodes []ocode.Ocode) []byte {
	var machineCode []byte
	sm := NewStackMachine()

	for _, ocode := range ocodes {
		code, err := processOcode(ocode, sm)
		if err != nil {

		}
		machineCode = append(machineCode, code...)
	}

	return machineCode
}

// processOcode processes a single ocode and returns the corresponding machine code.
func processOcode(oc ocode.Ocode, sm *StackMachine) ([]byte, error) {
	switch oc.Kind {
	case ocode.OpLN:
		return handleLN(oc.Operands, sm)
	case ocode.OpDB:
		return handleDB(oc.Operands), nil
	case ocode.OpDW:
		return handleDW(oc.Operands), nil
	case ocode.OpDD:
		return handleDD(oc.Operands), nil
	default:
		return handleNoParamOpcode(oc), nil
	}
}

// handleNoParamOpcode handles opcodes that do not require parameters.
func handleNoParamOpcode(ocode ocode.Ocode) []byte {
	if _, exists := opcodeMap[ocode.Kind]; exists {
		return GenerateX86NoParam(ocode)
	}
	return nil
}
