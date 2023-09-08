package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateTimeOff(w http.ResponseWriter, r *http.Request) error {
	timeOffRequest := &types.TimeOffRequest{}

	if err := json.NewDecoder(r.Body).Decode(&timeOffRequest); err != nil {
		return fmt.Errorf("error decoding timeOffRequest body: %v", err)
	}

	var timeOffReq *types.TimeOff

	newTimeOffRequest := timeOffReq.NewTimeOffRequest(timeOffRequest.StartDate, timeOffRequest.EndDate, timeOffRequest.Type)

	if err := s.store.CreateTimeOffRequest(newTimeOffRequest); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newTimeOffRequest)
}

func (s *APIServer) handleGetTimeOffRequests(w http.ResponseWriter, r *http.Request) error {
	requests, err := s.store.GetTimeOffRequests()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, requests)
}
