package types

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

func (*TimeOff) NewTimeOffRequest(start_date, end_date, time_off_type string) *TimeOff {
	return &TimeOff{
		Type:      time_off_type,
		StartDate: start_date,
		EndDate:   end_date,
		CreatedAt: time.Now().UTC(),
	}
}
