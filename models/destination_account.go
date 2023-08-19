package models

type DestinationAccount struct {
	Model

	URL           string      `json:"url"`
	AccountTypeID int         `json:"account_type_id"`
	AccountType   AccountType `json:"account_type"`
}
