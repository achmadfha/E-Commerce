package usersRepository

import (
	"E-Commerce/models/dto/usersDto"
	"E-Commerce/src/users"
	"database/sql"
	"github.com/google/uuid"
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
	WHERE
	  is_deleted = FALSE
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

func (u userRepository) CheckUserProfileExists(usrID string) (bool, error) {
	query := `SELECT
	  EXISTS(
		SELECT
		  1
		FROM
		  user_profiles
		WHERE
		  user_id = $1
	  )`

	var exists bool
	err := u.db.QueryRow(query, usrID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u userRepository) RetrieveUsersByID(usrID string) (usrData usersDto.UserResponse, err error) {
	query := `SELECT
	  u.user_id,
	  u.username,
	  u.email,
	  u.created_at,
	  u.updated_at,
	  up.full_name,
	  up.address,
	  up.city,
	  up.state,
	  up.country,
	  up.postal_code,
	  up.phone
	FROM
	  users u
	  JOIN user_profiles up ON u.user_id = up.user_id
	WHERE
	  u.user_id = $1`

	row := u.db.QueryRow(query, usrID)
	err = row.Scan(
		&usrData.UserID,
		&usrData.Username,
		&usrData.Email,
		&usrData.CreatedAt,
		&usrData.UpdatedAt,
		&usrData.UserProfile.FullName,
		&usrData.UserProfile.Address,
		&usrData.UserProfile.City,
		&usrData.UserProfile.State,
		&usrData.UserProfile.Country,
		&usrData.UserProfile.PostalCode,
		&usrData.UserProfile.Phone,
	)

	if err != nil {
		return usrData, err
	}

	return usrData, nil
}

func (u userRepository) UpdateProfiles(user usersDto.UserUpdate) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Check if the user profile exists
	var profileExists bool
	checkProfileQuery := `SELECT EXISTS(SELECT 1 FROM user_profiles WHERE user_id = $1)`
	err = tx.QueryRow(checkProfileQuery, user.UserID).Scan(&profileExists)
	if err != nil {
		return err
	}

	if profileExists {
		// Update the existing profile
		updateProfileQuery := `UPDATE
		  user_profiles
		SET
		  full_name = $2,
		  address = $3,
		  city = $4,
		  state = $5,
		  country = $6,
		  postal_code = $7,
		  phone = $8
		WHERE
		  user_id = $1`

		_, err = tx.Exec(updateProfileQuery, user.UserID, user.FullName, user.Address, user.City, user.State, user.Country, user.PostalCode, user.Phone)
		if err != nil {
			return err
		}
	} else {
		// Create a new profile
		userProfileID, err := uuid.NewRandom()
		insertProfileQuery := `INSERT INTO user_profiles (user_profile_id, user_id, full_name, address, city, state, country, postal_code, phone)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

		_, err = tx.Exec(insertProfileQuery, userProfileID, user.UserID, user.FullName, user.Address, user.City, user.State, user.Country, user.PostalCode, user.Phone)
		if err != nil {
			return err
		}
	}

	return nil
}
