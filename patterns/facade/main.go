package main

import "fmt"

// SubsystemA представляет первую подсистему.
type SubsystemA struct {
}

// Методы SubsystemA
func (s *SubsystemA) OperationA1() {
	fmt.Println("SubsystemA: OperationA1")
}

func (s *SubsystemA) OperationA2() {
	fmt.Println("SubsystemA: OperationA2")
}

// SubsystemB представляет вторую подсистему.
type SubsystemB struct {
}

// Методы SubsystemB
func (s *SubsystemB) OperationB1() {
	fmt.Println("SubsystemB: OperationB1")
}

func (s *SubsystemB) OperationB2() {
	fmt.Println("SubsystemB: OperationB2")
}

// Facade представляет собой фасадный объект, который предоставляет упрощенный интерфейс
// к сложной подсистеме, объединяя операции подсистем A и B.
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

// Инициализация фасада где создаются все структуры-подсистемы
func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

// Operation осуществляет упрощенный доступ к сложной системе через фасад.
func (f *Facade) Operation() {
	fmt.Println("Facade initializes subsystems:")
	f.subsystemA.OperationA1()
	f.subsystemB.OperationB1()
	fmt.Println("Facade orders subsystems to perform the action:")
	f.subsystemA.OperationA2()
	f.subsystemB.OperationB2()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
