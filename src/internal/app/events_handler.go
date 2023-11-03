package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent dto.CreateEventDTO
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if newEvent.EventName == "" {
		response.ErrorResponse(w, http.StatusBadRequest, messages.AllFieldMustBeFilled)
		return
	}

	eventModel := datastruct.Event{
		EventName: newEvent.EventName,
	}

	createdEvent, err := m.eventService.CreateEvent(eventModel)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToParseID)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataCreation, createdEvent)
}

func (m *MicroserviceServer) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventID := params["event_id"]

	parsedEventID, err := utils.ParseStrToUint(eventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	_, err = m.eventService.DeleteEvent(*parsedEventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToDeleteData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataDeletion, nil)
	return
}

func (m *MicroserviceServer) GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventID := params["event_id"]

	parsedEventID, err := utils.ParseStrToUint(eventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	event, err := m.eventService.GetEvent(*parsedEventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToGetData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, event)
	return
}

func (m *MicroserviceServer) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var updateEvent dto.UpdateEventDTO
	params := mux.Vars(r)
	eventID := params["event_id"]

	err := json.NewDecoder(r.Body).Decode(&updateEvent)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	parsedEventID, err := utils.ParseStrToUint(eventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	updatedEvent, err := m.eventService.UpdateEvent(*parsedEventID, updateEvent)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToUpdateData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataUpdate, updatedEvent)
	return
}
