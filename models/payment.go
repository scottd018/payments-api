package models

import (
	"time"
)

type Payment struct {
	Model

	Amount    float32    `json:"amount"`
	Due       *time.Time `json:"due"`
	Automatic bool       `json:"automatic"`

	SourceAccountID      int                `json:"source_account_id"`
	SourceAccount        SourceAccount      `json:"source_account"`
	DestinationAccountID int                `json:"destination_account_id"`
	DestinationAccount   DestinationAccount `json:"destination_account"`
	CategoryID           int                `json:"category_id"`
	Category             Category           `json:"category"`
	FrequencyID          int                `json:"frequency_id"`
	Frequency            Frequency          `json:"frequency"`
}
