package dto

type Quote struct {
	Author string `json:"author" binding:"required"`
	Text   string `json:"quote" binding:"required"`
}
