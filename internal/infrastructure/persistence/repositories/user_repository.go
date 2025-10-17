package repositories

import "database/sql"

type UserRepository struct {
	DB *sql.DB
}

const (
	got              = "%w: got %w"
	QueryGetAllUsers = `SELECT id, first_name, last_name, user_name, email, password, gender, birth_date, phone, country, language, information, profile_pic, web_site, visibility, created_at, updated_at, deleted_at
						FROM users`
)
