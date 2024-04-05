package models

import (
	"errors"

	"de.anikate/events-api/db"
	"de.anikate/events-api/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	res, err := db.DB.Exec(query, user.Email, password)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	user.ID = id

	return nil
}

func (u *User) ValidateUser() error {
	query := `SELECT password, id FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var hash string
	err := row.Scan(&hash, &u.ID)

	if err != nil {
		return errors.New("invalid credentials")
	}

	if !utils.CompareHashAndPassword(u.Password, hash) {
		return errors.New("invalid credentials")
	}

	return nil
}
