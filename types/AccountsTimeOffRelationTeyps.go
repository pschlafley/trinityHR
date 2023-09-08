package types

type AccountsTimeOffRelationRequest struct {
	AccountID int `json:"account_id"`
	TimeOffID int `json:"time_off_id"`
}

type AccountsTimeOffRelationTable struct {
	ID        int `json:"id"`
	AccountID int `json:"account_id"`
	TimeOffID int `json:"time_off_id"`
}

type AccountTimeOffRelationQueryData struct {
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (*AccountsTimeOffRelationTable) NewAccountsTimeOffRelationTable(account_id, time_off_id int) *AccountsTimeOffRelationTable {
	return &AccountsTimeOffRelationTable{
		AccountID: account_id,
		TimeOffID: time_off_id,
	}
}
