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
	Creator  string `sql:"index"`
}

func AllEvents(creator string) []Event {
	openDB()
	defer DB.Close()

	var events []Event
	DB.Where("creator = ?", creator).Find(&events)

	return events
}

func NewEvent(e *Event) {
	openDB()
	defer DB.Close()

	DB.Create(&e)
}

func FindEvent(id uint, creator string) (Event, error) {
	openDB()
	defer DB.Close()

	var e Event
	if DB.First(&e, id).RecordNotFound() {
		return Event{}, fmt.Errorf("ID not found: %d", uint(id))
	}

	return e, nil
}

func FindEvents(code string, creator string) []Event {
	openDB()
	defer DB.Close()

	var r []Event
	DB.Where("creator = ? AND code = ?", creator, code).Find(&r)

	return r
}

func DeleteEvent(e *Event) {
	openDB()
	defer DB.Close()

	DB.First(e, e.ID).Delete(e)
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
		Find(&event, uint(id)).
		Update(updates)

	return event, nil
}
