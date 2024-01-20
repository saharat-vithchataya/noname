package repository

type Account struct {
	ID             int    `db:"id"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	IsVerified     bool   `db:"is_verified"`
	OpeningDate    string `db:"opening_date"`
	DeletedAt      string `db:"deleted_at"`
}

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAccount(int) (*Account, error)
	GetAccountByEmail(string) (*Account, error)
}
