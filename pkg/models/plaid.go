package models

type PublicToken struct {
	PublicToken string `json:"publicToken"`
}

type PrivateToken struct {
	UserID       *string `json:"id"`
	PrivateToken *string `json:"privateToken"`
	ItemId       *string `json:"itemId"`
	IsNew        *bool   `json:"isNew"`
	Cursor       *string `json:"cursor"`
}

type LinkToken struct {
	LinkToken *string `json:"linkToken"`
}
