package models

type PublicToken struct {
	PublicToken string `json:"publicToken"`
}

type PrivateToken struct {
	UserID       *string `json:"id"`
	PrivateToken *string `json:"privateToken"`
	ItemId       *string `json:"itemId"`
}

type LinkToken struct {
	LinkToken *string `json:"linkToken"`
}
