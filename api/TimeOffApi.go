package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pschlafley/trinityHR/DbTypes"
)

func (s *APIServer) handleCreateTimeOff(w http.ResponseWriter, r *http.Request) error {
	timeOffRequest := &DbTypes.TimeOffRequest{}

	if err := json.NewDecoder(r.Body).Decode(&timeOffRequest); err != nil {
		return fmt.Errorf("error decoding timeOffRequest body: %v", err)
	}

	timeOffID, err := s.store.CreateTimeOffRequest(timeOffRequest)

	if err != nil {
		return err
	}

	var timeOffReq *DbTypes.TimeOff

	newTimeOffRequest := timeOffReq.NewTimeOffRequest(timeOffID, *timeOffRequest)

	return WriteJSON(w, http.StatusOK, newTimeOffRequest)
}

func (s *APIServer) handleGetTimeOffRequests(w http.ResponseWriter, r *http.Request) error {
	requests, err := s.store.GetTimeOffRequests()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, requests)
}
