package request

type Transaction struct {
	UserID      int    `json:"user_id" validate:"required"`
	Amount      int    `json:"amount" validate:"required"`
	Description string `json:"description" validate:"required"`
}
