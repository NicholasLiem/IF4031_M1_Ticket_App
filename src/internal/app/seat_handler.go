package app

import (
	"encoding/json"
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
	"github.com/gorilla/mux"
)

func (m *MicroserviceServer) CreateSeat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventID := params["event_id"]

	var newSeat dto.CreateSeatDTO
	err := json.NewDecoder(r.Body).Decode(&newSeat)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	if newSeat.Status == "" || !datastruct.IsValidStatus(newSeat.Status) {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	parsedEventID, err := utils.ParseStrToUint(eventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	seatModel := datastruct.Seat{
		Status:  newSeat.Status,
		EventID: *parsedEventID,
	}

	createdSeat, err := m.seatService.CreateSeat(seatModel, *parsedEventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToCreateData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataCreation, createdSeat)
}

func (m *MicroserviceServer) UpdateSeat(w http.ResponseWriter, r *http.Request) {
	var updateSeat dto.UpdateSeatDTO
	params := mux.Vars(r)
	seatID := params["seat_id"]

	err := json.NewDecoder(r.Body).Decode(&updateSeat)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	parsedSeatID, err := utils.ParseStrToUint(seatID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	existingSeat, err := m.seatService.GetSeat(*parsedSeatID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.DataNotFound)
		return
	}

	updatedSeatData := datastruct.Seat{
		Status:  updateSeat.Status,
		EventID: existingSeat.EventID,
	}

	updatedSeat, err := m.seatService.UpdateSeat(*parsedSeatID, updatedSeatData)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToUpdateData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataUpdate, updatedSeat)
}

func (m *MicroserviceServer) DeleteSeat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	seatID := params["seat_id"]

	parsedSeatID, err := utils.ParseStrToUint(seatID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	_, err = m.seatService.GetSeat(*parsedSeatID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.DataNotFound)
		return
	}

	_, err = m.seatService.DeleteSeat(*parsedSeatID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToDeleteData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataDeletion, nil)
}

func (m *MicroserviceServer) GetSeat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	seatID := params["seat_id"]

	parsedSeatID, err := utils.ParseStrToUint(seatID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	seat, err := m.seatService.GetSeat(*parsedSeatID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToGetData)
		return
	}

	if seat == nil {
		response.ErrorResponse(w, http.StatusNotFound, messages.DataNotFound)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, seat)
}

func (m *MicroserviceServer) GetSeatsForEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventID := params["event_id"]

	parsedEventID, err := utils.ParseStrToUint(eventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.FailToParseID)
		return
	}

	event, httpError := m.eventService.GetEvent(*parsedEventID)
	if httpError != nil || event == nil {
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
		return
	}

	seats, err := m.seatService.GetSeatsForEvent(*parsedEventID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToGetData)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, seats)
}
