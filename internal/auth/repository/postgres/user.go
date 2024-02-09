package postgres

import (
	"database/sql"
	"go-clean-architecture/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(user *models.User, roleID int) error {
	query := "INSERT INTO users(name, surname, username, password, role_id) VALUES ($1, $2, $3, $4, $5)"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.Name,
		user.Surname,
		user.Username,
		user.Password,
		roleID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	query := "SELECT u.id, u.name, u.surname, u.username, u.password, r.role_name " +
		"FROM users u INNER JOIN roles r ON u.role_id = r.id " +
		"WHERE u.id = $1 LIMIT 1"

	row := r.db.QueryRow(query, userID)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT u.id, u.name, u.surname, u.username, u.password, r.role_name " +
		"FROM users u INNER JOIN roles r ON u.role_id = r.id " +
		"WHERE u.username = $1 LIMIT 1"

	row := r.db.QueryRow(query, username)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) GetAllUsers() (*[]models.User, error) {
	query := "SELECT u.id, u.name, u.surname, u.username, u.password, r.role_name " +
		"FROM users u INNER JOIN roles r ON u.role_id = r.id"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Username,
			&user.Password,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

func (r UserRepository) DeleteUserByID(userID int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) UpdateUserByID(userID int, user *models.User) error {
	query := "UPDATE users SET name=$1, surname=$2, username=$3, role_id=$4 WHERE id=$5"
	_, err := r.db.Exec(query, user.Name, user.Surname, user.Username, user.RoleID, userID)
	if err != nil {
		return err
	}
	return nil
}
