package API

import (
	"fmt"
	"net/http"

	"../model"
)

func AllEvents(w http.ResponseWriter, r *http.Request) {
	events := model.AllEvents()
	send(
		map[string]interface{}{"events": events}, true, "Data found", w)
}

func NewEvent(w http.ResponseWriter, r *http.Request) {
	event := &model.Event{}
	decode(event, w, r)
	model.NewEvent(event)

	send(
		map[string]interface{}{"event": event},
		true,
		fmt.Sprintln("Event", event.Code, "has been created successfully"),
		w)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	event := &model.Event{}
	decode(event, w, r)
	model.DeleteEvent(event)

	send(map[string]interface{}{"event": event}, true, "Event Deleted", w)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	e := model.Event{}
	decode(&e, w, r)
	model.UpdateEvent(&e)

	send(map[string]interface{}{"event": e}, true, "Event Updated", w)
}
