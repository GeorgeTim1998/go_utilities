package main

import "fmt"

// State интерфейс для состояний принтера
type State interface {
	Print()
	AddPaper()
	Service()
}

// Printer контекст, который использует состояния
type Printer struct {
	state State
}

// NewPrinter создает новый принтер в начальном состоянии
func NewPrinter() *Printer {
	return &Printer{state: &ReadyState{}}
}

// SetState позволяет изменить текущее состояние принтера
func (p *Printer) SetState(state State) {
	p.state = state
}

// Print вызывает метод Print текущего состояния
func (p *Printer) Print() {
	p.state.Print()
}

// AddPaper вызывает метод AddPaper текущего состояния
func (p *Printer) AddPaper() {
	p.state.AddPaper()
}

// Service вызывает метод Service текущего состояния
func (p *Printer) Service() {
	p.state.Service()
}

// ReadyState состояние, когда принтер готов к печати
type ReadyState struct{}

func (r *ReadyState) Print() {
	fmt.Println("Печать началась.")
	// Логика перехода в состояние "Печатает"
}

func (r *ReadyState) AddPaper() {
	fmt.Println("Бумага уже есть.")
}

func (r *ReadyState) Service() {
	fmt.Println("Принтер не нуждается в обслуживании.")
}

// PrintingState состояние, когда принтер печатает
type PrintingState struct{}

func (p *PrintingState) Print() {
	fmt.Println("Принтер уже печатает.")
}

func (p *PrintingState) AddPaper() {
	fmt.Println("Принтер печатает. Добавление бумаги невозможно.")
}

func (p *PrintingState) Service() {
	fmt.Println("Принтер печатает. Обслуживание невозможно.")
}

// NoPaperState состояние, когда принтер нуждается в бумаге
type NoPaperState struct{}

func (n *NoPaperState) Print() {
	fmt.Println("Нет бумаги. Печать невозможна.")
}

func (n *NoPaperState) AddPaper() {
	fmt.Println("Бумага добавлена. Принтер готов к печати.")
	// Логика перехода в состояние "Готов к печати"
}

func (n *NoPaperState) Service() {
	fmt.Println("Принтер не нуждается в обслуживании.")
}

// ServiceState состояние, когда принтер нуждается в обслуживании
type ServiceState struct{}

func (s *ServiceState) Print() {
	fmt.Println("Принтер нуждается в обслуживании. Печать невозможна.")
}

func (s *ServiceState) AddPaper() {
	fmt.Println("Принтер нуждается в обслуживании. Добавление бумаги невозможно.")
}

func (s *ServiceState) Service() {
	fmt.Println("Принтер обслуживается.")
	// Логика перехода в состояние "Готов к печати"
}

func main() {
	printer := NewPrinter()

	printer.Print()    // Вывод: Печать началась.
	printer.AddPaper() // Вывод: Бумага уже есть.
	printer.Service()  // Вывод: Принтер не нуждается в обслуживании.

	printer.SetState(&NoPaperState{})
	printer.Print()    // Вывод: Нет бумаги. Печать невозможна.
	printer.AddPaper() // Вывод: Бумага добавлена. Принтер готов к печати.

	printer.SetState(&ServiceState{})
	printer.Print()   // Вывод: Принтер нуждается в обслуживании. Печать невозможна.
	printer.Service() // Вывод: Принтер обслуживается.
}
