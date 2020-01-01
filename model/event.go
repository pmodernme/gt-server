package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Event - an individual item event, scanning in or out
type Event struct {
	gorm.Model
	Code     string `json:"code" sql:"index"`
	Incoming bool   `json:"incoming"`
}

func AllEvents() []Event {
	openDB()
	defer DB.Close()

	var events []Event
	DB.Find(&events)

	return events
}

func NewEvent(e *Event) {
	openDB()
	defer DB.Close()

	DB.Create(&e)
}

func DeleteEvent(e *Event) {
	openDB()
	defer DB.Close()

	DB.Where("ID = ?", e.ID).Find(e).Delete(e)
}

func UpdateEvent(updates map[string]interface{}) (Event, error) {
	openDB()
	defer DB.Close()

	id, ok := updates["ID"].(float64)
	if !ok {
		return Event{}, fmt.Errorf("Invalid ID: %d", uint(id))
	}

	var event Event
	DB.Unscoped().
		Where("ID = ?", uint(id)).
		Find(&event).
		Update(updates)

	return event, nil
}
