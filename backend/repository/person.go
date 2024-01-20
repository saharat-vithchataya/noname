package repository

import "time"

type Person struct {
	UserID    int       `db:"user_id"`
	Firstname string    `db:"firstname"`
	Lastname  string    `db:"lastname"`
	Birthday  time.Time `db:"birthday"`
	Country   string    `db:"country"`
	Gender    string    `db:"gender"`
	Phone     string    `db:"phone"`
}

type PersonRepository interface {
	Create(Person) (*Person, error)
	GetPerson(int) (*Person, error)
}
