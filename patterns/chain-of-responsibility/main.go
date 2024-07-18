package main

import (
	"fmt"
)

// Общий интерфейс для всех проверок в цепочке
type Handler interface {
	SetNext(handler Handler)      // Установка следующего обработчика
	HandleRequest(request string) // Обработка запроса
}

// Базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler // Следующий обработчик в цепочке
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.nextHandler = handler
}

// Специфическая проверка №1
type Check1Handler struct {
	BaseHandler
}

func (h *Check1Handler) HandleRequest(request string) {
	if request == "check1" {
		fmt.Println("Check1Handler: handling request")
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

// Специфическая проверка №2
type Check2Handler struct {
	BaseHandler
}

func (h *Check2Handler) HandleRequest(request string) {
	if request == "check2" {
		fmt.Println("Check2Handler: handling request")
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

// Специфическая проверка №3
type Check3Handler struct {
	BaseHandler
}

func (h *Check3Handler) HandleRequest(request string) {
	if request == "check3" {
		fmt.Println("Check3Handler: handling request")
	} else if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

func main() {
	// Создание объектов обработчиков
	handler1 := &Check1Handler{}
	handler2 := &Check2Handler{}
	handler3 := &Check3Handler{}

	// Установка цепочки обработчиков
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Тестирование работы цепочки
	requests := []string{"check1", "check2", "check3", "unknown"}

	for _, req := range requests {
		fmt.Printf("Sending request '%s' through the chain:\n", req)
		handler1.HandleRequest(req)
		fmt.Println("-----------------------------")
	}
}
