package DbTypes

import "time"

type TimeOffRequest struct {
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type TimeOff struct {
	TimeOffID int       `json:"time_off_id"`
	Type      string    `json:"type"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

func (*TimeOff) NewTimeOffRequest(timeOffID int, request TimeOffRequest) *TimeOff {
	return &TimeOff{
		TimeOffID: timeOffID,
		Type:      request.Type,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
		CreatedAt: time.Now().UTC(),
	}
}
