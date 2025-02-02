package data_model

type User struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	ApiKey string `db:"api_key"`
}
