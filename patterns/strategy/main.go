package main

import (
	"fmt"
)

// Интерфейс стратегии
type RouteStrategy interface {
	CalculateRoute(start, end string) string
}

// Конкретная стратегия для автомобильных маршрутов
type CarRouteStrategy struct{}

func (c *CarRouteStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Calculating car route from %s to %s", start, end)
}

// Конкретная стратегия для пеших маршрутов
type WalkingRouteStrategy struct{}

func (w *WalkingRouteStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Calculating walking route from %s to %s", start, end)
}

// Конкретная стратегия для маршрутов общественного транспорта
type PublicTransportRouteStrategy struct{}

func (p *PublicTransportRouteStrategy) CalculateRoute(start, end string) string {
	return fmt.Sprintf("Calculating public transport route from %s to %s", start, end)
}

// Контекст, использующий стратегии
type Navigator struct {
	strategy RouteStrategy
}

// Метод для установки стратегии
func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

// Метод для расчета маршрута с использованием установленной стратегии
func (n *Navigator) CalculateRoute(start, end string) {
	if n.strategy == nil {
		fmt.Println("No route strategy set")
		return
	}
	route := n.strategy.CalculateRoute(start, end)
	fmt.Println(route)
}

func main() {
	// Создание экземпляра навигатора
	navigator := &Navigator{}

	// Установка и использование стратегии автомобильного маршрута
	navigator.SetStrategy(&CarRouteStrategy{})
	navigator.CalculateRoute("Point A", "Point B")

	// Установка и использование стратегии пешего маршрута
	navigator.SetStrategy(&WalkingRouteStrategy{})
	navigator.CalculateRoute("Point A", "Point B")

	// Установка и использование стратегии маршрута общественного транспорта
	navigator.SetStrategy(&PublicTransportRouteStrategy{})
	navigator.CalculateRoute("Point A", "Point B")
}
