package main

import "fmt"

// PaymentMethod интерфейс для различных способов оплаты
type PaymentMethod interface {
	Pay(amount float64) string
}

// CreditCard структура для оплаты кредитной картой
type CreditCard struct{}

func (cc *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f using Credit Card", amount)
}

// PayPal структура для оплаты через PayPal
type PayPal struct{}

func (pp *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f using PayPal", amount)
}

// PaymentMethodFactory интерфейс для фабрики способов оплаты
type PaymentMethodFactory interface {
	CreatePaymentMethod() PaymentMethod
}

// CreditCardFactory структура фабрики для создания объектов CreditCard
type CreditCardFactory struct{}

func (ccf *CreditCardFactory) CreatePaymentMethod() PaymentMethod {
	return &CreditCard{}
}

// PayPalFactory структура фабрики для создания объектов PayPal
type PayPalFactory struct{}

func (ppf *PayPalFactory) CreatePaymentMethod() PaymentMethod {
	return &PayPal{}
}

func main() {
	var factory PaymentMethodFactory

	// Создание объекта CreditCard через фабрику
	factory = &CreditCardFactory{}
	paymentMethod := factory.CreatePaymentMethod()
	fmt.Println(paymentMethod.Pay(100.50))

	// Создание объекта PayPal через фабрику
	factory = &PayPalFactory{}
	paymentMethod = factory.CreatePaymentMethod()
	fmt.Println(paymentMethod.Pay(200.75))
}
