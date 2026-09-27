package models

import (
	"errors"
	"time"

	"Go-gin-rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func (e *Event) UpdateEventById(id int64) error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, datetime = ? WHERE id = ?`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("event not updated or event not found")
	}

	e.ID = id
	return nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, datetime, user_id FROM events WHERE id = ?`
	row := db.Db.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, datetime, user_id FROM events`
	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (e *Event) DeleteEventById() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("event not deleted or event not found")
	}

	return nil
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO registrations (event_id,user_id) VALUES (?, ?)`
	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE user_id = ? and event_id = ?`
	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, e.ID)
	if err != nil {
		return err
	}
	return nil
}
