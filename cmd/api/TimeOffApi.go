package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pschlafley/trinityHR/types"
)

func (s *APIServer) handleCreateTimeOff(c echo.Context) error {
	timeOffRequest := &types.TimeOffRequest{}

	if err := json.NewDecoder(c.Request().Body).Decode(&timeOffRequest); err != nil {
		return fmt.Errorf("error decoding timeOffRequest body: %v", err)
	}

	timeOffID, err := s.store.CreateTimeOffRequest(timeOffRequest)

	if err != nil {
		return err
	}

	var timeOffReq *types.TimeOff

	newTimeOffRequest := timeOffReq.NewTimeOffRequest(timeOffID, *timeOffRequest)

	return c.JSON(http.StatusOK, newTimeOffRequest)
}

func (s *APIServer) handleGetTimeOffRequests(c echo.Context) error {
	requests, err := s.store.GetTimeOffRequests()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, requests)
}
