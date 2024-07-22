package main

import (
	"errors"
	"time"
)

type Event struct {
	ID     string    `json:"id"`
	UserID int       `json:"user_id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
}

var calendar = make(map[string]Event)

func createEvent(event Event) error {
	if _, exists := calendar[event.ID]; exists {
		return errors.New("событие уже существует")
	}
	calendar[event.ID] = event
	return nil
}

func updateEvent(event Event) error {
	if _, exists := calendar[event.ID]; !exists {
		return errors.New("событие не найдено")
	}
	calendar[event.ID] = event
	return nil
}

func deleteEvent(eventID string) error {
	if _, exists := calendar[eventID]; !exists {
		return errors.New("событие не найдено")
	}
	delete(calendar, eventID)
	return nil
}

func getEventsForDay(date time.Time) ([]Event, error) {
	var events []Event
	for _, event := range calendar {
		if event.Date.Year() == date.Year() && event.Date.YearDay() == date.YearDay() {
			events = append(events, event)
		}
	}
	return events, nil
}

func getEventsForWeek(date time.Time) ([]Event, error) {
	var events []Event
	year, week := date.ISOWeek()
	for _, event := range calendar {
		eventYear, eventWeek := event.Date.ISOWeek()
		if eventYear == year && eventWeek == week {
			events = append(events, event)
		}
	}
	return events, nil
}

func getEventsForMonth(date time.Time) ([]Event, error) {
	var events []Event
	for _, event := range calendar {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			events = append(events, event)
		}
	}
	return events, nil
}
