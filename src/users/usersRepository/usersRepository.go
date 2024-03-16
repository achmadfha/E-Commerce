package usersRepository

import (
	"E-Commerce/models/dto/usersDto"
	"E-Commerce/src/users"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) users.UserRepository {
	return userRepository{db}
}

func (u userRepository) RetrieveAllUsers(page, pageSize int) ([]usersDto.User, error) {
	offset := (page - 1) * pageSize
	limit := pageSize

	query := `SELECT
	  user_id,
	  username,
	  password,
	  email,
	  role,
	  created_at,
	  updated_at
	FROM
	  users
	LIMIT $1 OFFSET $2`

	rows, err := u.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []usersDto.User
	for rows.Next() {
		var user usersDto.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepository) CountAllUsers() (int, error) {
	var count int

	query := `SELECT COUNT(*) FROM users`
	rows := u.db.QueryRow(query)
	err := rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
