package models

import (
	"time"

	"de.anikate/events-api/db"
)

type Event struct {
	ID          int64
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	AuthorID    int64
}

func (event *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, date, author_id)
		VALUES (?, ?, ?, ?, ?)
	`

	res, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.Date, event.AuthorID)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	event.ID = id

	return nil
}

func GetAllEvents() ([]*Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []*Event
	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.AuthorID)

		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(query, id)

	event := Event{}
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.AuthorID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, date = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.Date, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (event *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	_, err := db.DB.Exec(query, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Register(userID int64) error {
	query := `
		INSERT INTO registrations (event_id, user_id)
		VALUES (?, ?)
	`

	_, err := db.DB.Exec(query, e.ID, userID)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Unregister(userID int64) error {
	query := `
		DELETE FROM registrations
		WHERE event_id = ? AND user_id = ?
	`

	_, err := db.DB.Exec(query, e.ID, userID)

	if err != nil {
		return err
	}

	return nil
}
