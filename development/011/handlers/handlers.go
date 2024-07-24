package handlers

import (
	"011/calendar"
	"011/utils"
	"net/http"
	"time"
)

// Обработчик для создания события
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event calendar.Event
	err := utils.ParseEvent(r, &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = calendar.CreateEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"result": "событие создано"})
}

// Обработчик для обновления события
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event calendar.Event
	err := utils.ParseEvent(r, &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = utils.ParseEvent(r, &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"result": "событие обновлено"})
}

// Обработчик для удаления события
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// парсит параметры, переданные в url запроса
	eventID := r.URL.Query().Get("id")
	if eventID == "" {
		http.Error(w, "id не указан", http.StatusBadRequest)
		return
	}
	err := calendar.DeleteEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"result": "событие удалено"})
}

// Обработчик для получения событий за день
func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "неправильный формат даты", http.StatusBadRequest)
		return
	}
	events, err := calendar.GetEventsForDay(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	utils.RespondJSON(w, http.StatusOK, events)
}

// Обработчик для получения событий за неделю
func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "неправильный формат даты", http.StatusBadRequest)
		return
	}
	events, err := calendar.GetEventsForWeek(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	utils.RespondJSON(w, http.StatusOK, events)
}

// Обработчик для получения событий за месяц
func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "неправильный формат даты", http.StatusBadRequest)
		return
	}
	events, err := calendar.GetEventsForMonth(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	utils.RespondJSON(w, http.StatusOK, events)
}
