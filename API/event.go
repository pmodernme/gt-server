package API

import (
	"fmt"
	"net/http"
	"strconv"

	"../auth"
	"../model"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if _, ok := q["id"]; ok {
		FindEvent(w, r)
	} else if _, ok := q["code"]; ok {
		FindEvents(w, r)
	} else {
		AllEvents(w, r)
	}
}

func AllEvents(w http.ResponseWriter, r *http.Request) {
	username, err := auth.GetUsername(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid Token")
		return
	}

	events := model.AllEvents(username)
	send(
		map[string]interface{}{"events": events},
		true,
		"Data found",
		w)
}

func FindEvent(w http.ResponseWriter, r *http.Request) {
	username, err := auth.GetUsername(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid Token")
		return
	}

	q := r.URL.Query()
	sid, ok := q["id"]
	if !ok {
		writeError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	id, err := strconv.ParseUint(sid[0], 10, 8)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid Query Parameters")
		return
	}

	event, err := model.FindEvent(uint(id), username)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Could not find event with ID: %d", id))
		return
	}

	send(
		map[string]interface{}{"event": event},
		true,
		fmt.Sprint("Event", event.ID, "has been found successfully"),
		w)
}

func FindEvents(w http.ResponseWriter, r *http.Request) {
	username, err := auth.GetUsername(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid Token")
		return
	}

	q := r.URL.Query()
	scode, ok := q["code"]
	if !ok {
		writeError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	code := scode[0]

	result := model.FindEvents(code, username)

	var message string
	if len(result) > 0 {
		message = fmt.Sprint("Events for code", code, "have been found successfully")
	} else {
		message = fmt.Sprint("No Events found for code: ", code)
	}

	send(
		map[string]interface{}{"events": result},
		true,
		message,
		w)
}

func NewEvent(w http.ResponseWriter, r *http.Request) {
	username, err := auth.GetUsername(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid Token")
		return
	}

	event := &model.Event{}
	decode(event, w, r)
	event.Creator = username

	model.NewEvent(event)

	send(
		map[string]interface{}{"event": event},
		true,
		fmt.Sprint("Event", event.Code, "has been created successfully"),
		w)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	username, err := auth.GetUsername(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid Token")
		return
	}

	event := &model.Event{}
	decode(event, w, r)
	event.Creator = username
	model.DeleteEvent(event)

	send(map[string]interface{}{"event": event}, true, "Event Deleted", w)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	updates := map[string]interface{}{}
	decode(&updates, w, r)
	e, err := model.UpdateEvent(updates)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	send(map[string]interface{}{"event": e}, true, "Event Updated", w)
}
