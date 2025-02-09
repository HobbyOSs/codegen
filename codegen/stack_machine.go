package backend

import "fmt"

type StackMachine struct {
	stack     []DWord // DWord 型のスタック
	registers map[string]DWord
}

func NewStackMachine() *StackMachine {
	return &StackMachine{
		stack:     []DWord{},
		registers: make(map[string]DWord),
	}
}

// Push スタックに値を積む
func (s *StackMachine) Push(value DWord) {
	s.stack = append(s.stack, value)
}

// Pop スタックから値を取り出す
func (s *StackMachine) Pop() (DWord, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("stack underflow")
	}
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value, nil
}

// LoadRegister レジスタの値をスタックに積む
func (s *StackMachine) LoadRegister(reg string) error {
	value, exists := s.registers[reg]
	if !exists {
		return fmt.Errorf("unknown register: %s", reg)
	}
	s.Push(value)
	return nil
}

// StoreRegister スタックトップをレジスタに保存
func (s *StackMachine) StoreRegister(reg string) error {
	value, err := s.Pop()
	if err != nil {
		return err
	}
	s.registers[reg] = value
	return nil
}
