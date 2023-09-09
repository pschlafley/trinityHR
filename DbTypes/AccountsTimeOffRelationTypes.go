package DbTypes

type AccountsTimeOffRelationRequest struct {
	AccountsTimeOffRelationID int `json:"accounts_timeOff_relation_id"`
	AccountID                 int `json:"account_id"`
	TimeOffID                 int `json:"time_off_id"`
}

type AccountsTimeOffRelationTable struct {
	AccountsTimeOffRelationID int `json:"accounts_timeOff_relation_id"`
	AccountID                 int `json:"account_id"`
	TimeOffID                 int `json:"time_off_id"`
}

type AccountTimeOffRelationQueryData struct {
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Type      string `json:"type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (*AccountsTimeOffRelationTable) NewAccountsTimeOffRelationTable(accountsTimeOffRelationID, account_id, time_off_id int) *AccountsTimeOffRelationTable {
	return &AccountsTimeOffRelationTable{
		AccountsTimeOffRelationID: accountsTimeOffRelationID,
		AccountID:                 account_id,
		TimeOffID:                 time_off_id,
	}
}
