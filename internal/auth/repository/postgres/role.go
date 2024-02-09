package postgres

import (
	"database/sql"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r RoleRepository) CreateRole(role string) error {
	query := "INSERT INTO roles(role_name) VALUES ($1)"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(role)
	if err != nil {
		return err
	}

	return nil
}

func (r RoleRepository) GetNameByID(roleID int) (string, error) {
	var roleName string
	query := "SELECT role_name FROM roles WHERE id = $1"

	row := r.db.QueryRow(query, roleID)
	err := row.Scan(&roleName)
	if err != nil {
		return "", err
	}

	return roleName, nil
}

func (r RoleRepository) GetIDByName(roleName string) (int, error) {
	var id int
	query := "SELECT id FROM roles WHERE role_name = $1"

	row := r.db.QueryRow(query, roleName)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
