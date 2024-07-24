package main

import (
	"011/handlers"
	"011/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", handlers.CreateEventHandler)
	mux.HandleFunc("/update_event", handlers.UpdateEventHandler)
	mux.HandleFunc("/delete_event", handlers.DeleteEventHandler)
	mux.HandleFunc("/events_for_day", handlers.EventsForDayHandler)
	mux.HandleFunc("/events_for_week", handlers.EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", handlers.EventsForMonthHandler)

	// Применяем middleware для логирования
	handler := middleware.LoggingMiddleware(mux)

	// Запуск HTTP-сервера на порту 8080
	log.Println("Запуск сервера на порту 8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
