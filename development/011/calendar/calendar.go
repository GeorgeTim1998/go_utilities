package calendar

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

func CreateEvent(event Event) error {
	if _, exists := calendar[event.ID]; exists {
		return errors.New("событие уже существует")
	}
	calendar[event.ID] = event
	return nil
}

func UpdateEvent(event Event) error {
	if _, exists := calendar[event.ID]; !exists {
		return errors.New("событие не найдено")
	}
	calendar[event.ID] = event
	return nil
}

func DeleteEvent(eventID string) error {
	if _, exists := calendar[eventID]; !exists {
		return errors.New("событие не найдено")
	}
	delete(calendar, eventID)
	return nil
}

func GetEventsForDay(date time.Time) ([]Event, error) {
	var events []Event
	for _, event := range calendar {
		if event.Date.Year() == date.Year() && event.Date.YearDay() == date.YearDay() {
			events = append(events, event)
		}
	}
	return events, nil
}

func GetEventsForWeek(date time.Time) ([]Event, error) {
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

func GetEventsForMonth(date time.Time) ([]Event, error) {
	var events []Event
	for _, event := range calendar {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			events = append(events, event)
		}
	}
	return events, nil
}
