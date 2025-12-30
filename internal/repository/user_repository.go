package repository

import (
	"database/sql"
	"github.com/andreybrumatti/crud-init/internal/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAllUsers() ([]entity.User, error) {
	query := `SELECT 
					id,
					first_name, 
					last_name, 
					email, 
					password 
			  FROM users`

	rows, err := ur.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
