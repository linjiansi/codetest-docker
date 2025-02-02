package model

type Product struct {
	userId      int
	amount      int
	description string
}

func NewProduct(userId, amount int, description string) *Product {
	return &Product{
		userId:      userId,
		amount:      amount,
		description: description,
	}
}

func (p *Product) UserId() int {
	return p.userId
}

func (p *Product) Amount() int {
	return p.amount
}

func (p *Product) Description() string {
	return p.description
}
