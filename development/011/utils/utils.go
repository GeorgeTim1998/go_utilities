package utils

import (
	"011/calendar"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

// Функция для парсинга и валидации параметров события
func ParseEvent(r *http.Request, event *calendar.Event) error {
	// парсит форму переданную в теле POST запроса
	err := r.ParseForm()
	if err != nil {
		return errors.New("не удалось распарсить форму")
	}
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return errors.New("неправильный формат user_id")
	}
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		return errors.New("неправильный формат даты")
	}
	event.UserID = userID
	event.Title = r.FormValue("title")
	event.Date = date
	event.ID = r.FormValue("id")
	return nil
}

// Функция для отправки JSON-ответа
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "не удалось сериализовать ответ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
