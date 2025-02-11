package x86_gen

import (
    "fmt"
    "github.com/HobbyOSs/codegen/variantstack"
)

type StackMachine struct {
	stack variantstack.VariantStack // VariantStack 型のスタック
}

func NewStackMachine() *StackMachine {
	return &StackMachine{
		stack: variantstack.VariantStack{},
	}
}

// Push スタックに値を積む
func (s *StackMachine) Push(value DWord) {
	s.stack.Push(value)
}

// Pop スタックから値を取り出す
func (s *StackMachine) Pop() (DWord, error) {
	item, ok := s.stack.Pop()
	if !ok {
		return 0, fmt.Errorf("stack underflow")
	}
	value, ok := item.(DWord)
	if !ok {
		return 0, fmt.Errorf("popped item is not of type DWord")
	}
	return value, nil
}

