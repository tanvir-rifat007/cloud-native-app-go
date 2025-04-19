package data

import "database/sql"

type Model struct {
	Newsletters NewsletterModel

}


func NewModel(db *sql.DB) Model {
	return Model{
		Newsletters: NewsletterModel{DB: db},
	}
}