package main

import "fmt"

// Node интерфейс для узлов дерева
type Node interface {
	Accept(Visitor)
}

// Visitor интерфейс для посетителей
type Visitor interface {
	VisitConcreteNodeA(*ConcreteNodeA)
	VisitConcreteNodeB(*ConcreteNodeB)
}

// ConcreteNodeA структура узла A
type ConcreteNodeA struct {
	Value string
}

// Accept метод для принятия посетителя
func (n *ConcreteNodeA) Accept(visitor Visitor) {
	visitor.VisitConcreteNodeA(n)
}

// ConcreteNodeB структура узла B
type ConcreteNodeB struct {
	Value int
}

// Accept метод для принятия посетителя
func (n *ConcreteNodeB) Accept(visitor Visitor) {
	visitor.VisitConcreteNodeB(n)
}

// ConcreteVisitor структура конкретного посетителя
type ConcreteVisitor struct{}

// VisitConcreteNodeA реализация метода для узла A
func (v *ConcreteVisitor) VisitConcreteNodeA(node *ConcreteNodeA) {
	fmt.Println("Visiting ConcreteNodeA with value:", node.Value)
}

// VisitConcreteNodeB реализация метода для узла B
func (v *ConcreteVisitor) VisitConcreteNodeB(node *ConcreteNodeB) {
	fmt.Println("Visiting ConcreteNodeB with value:", node.Value)
}

func main() {
	// Создание узлов
	nodeA := &ConcreteNodeA{Value: "example"}
	nodeB := &ConcreteNodeB{Value: 42}

	// Создание посетителя
	visitor := &ConcreteVisitor{}

	// Принятие посетителя узлами
	nodeA.Accept(visitor)
	nodeB.Accept(visitor)
}
