package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", createEventHandler)
	mux.HandleFunc("/update_event", updateEventHandler)
	mux.HandleFunc("/delete_event", deleteEventHandler)
	mux.HandleFunc("/events_for_day", eventsForDayHandler)
	mux.HandleFunc("/events_for_week", eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", eventsForMonthHandler)

	// Применяем middleware для логирования
	handler := loggingMiddleware(mux)

	// Запуск HTTP-сервера на порту 8080
	log.Println("Запуск сервера на порту 8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
