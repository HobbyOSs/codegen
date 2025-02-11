package x86_gen

import "fmt"

func handleLN(args []interface{}, sm *StackMachine) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("LN requires 1 argument")
	}
	value, ok := args[0].(DWord)
	if !ok {
		return nil, fmt.Errorf("argument is not of type Dword")
	}
	sm.Push(value)
	return nil, nil
}

func handleLL(args []interface{}, sm *StackMachine) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("LN requires 1 argument")
	}
	value, ok := args[0].(DWord)
	if !ok {
		return nil, fmt.Errorf("argument is not of type Dword")
	}
	sm.Push(value)
	return nil, nil
}
