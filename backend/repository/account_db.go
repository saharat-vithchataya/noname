package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (repo accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := "insert into accounts(email, hashed_password, is_verified, opening_date) values($1,$2,$3,$4) RETURNING id"
	// repo
	row := repo.db.QueryRow(query, acc.Email, acc.HashedPassword, acc.IsVerified, acc.OpeningDate)
	err := row.Scan(&acc.ID)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}
func (repo accountRepositoryDB) GetAccount(id int) (*Account, error) {

	query := "select id, email, hashed_password, is_verified, opening_date from accounts where id=$1"
	account := Account{}

	err := repo.db.Get(&account, query, id)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repo accountRepositoryDB) GetAccountByEmail(email string) (*Account, error) {

	query := "select id, email, hashed_password, is_verified, opening_date from accounts where email=$1"
	account := Account{}

	err := repo.db.Get(&account, query, email)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
