package main

import (
	"net/http"
	"time"
)

// Обработчик для создания события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := parseEvent(r, &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = createEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"result": "событие создано"})
}

// Обработчик для обновления события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := parseEvent(r, &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = updateEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"result": "событие обновлено"})
}

// Обработчик для удаления события
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Query().Get("id")
	if eventID == "" {
		http.Error(w, "id не указан", http.StatusBadRequest)
		return
	}
	err := deleteEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"result": "событие удалено"})
}

// Обработчик для получения событий за день
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "неправильный формат даты", http.StatusBadRequest)
		return
	}
	events, err := getEventsForDay(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	respondJSON(w, http.StatusOK, events)
}

// Обработчик для получения событий за неделю
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "неправильный формат даты", http.StatusBadRequest)
		return
	}
	events, err := getEventsForWeek(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	respondJSON(w, http.StatusOK, events)
}

// Обработчик для получения событий за месяц
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "неправильный формат даты", http.StatusBadRequest)
		return
	}
	events, err := getEventsForMonth(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	respondJSON(w, http.StatusOK, events)
}
