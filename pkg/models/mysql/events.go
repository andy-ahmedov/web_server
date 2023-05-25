package mysql

import (
	"database/sql"
	"errors"

	"github.com/andy-ahmedov/web_server/pkg/models"
)

type EventModel struct {
	DB *sql.DB
}

func (m *EventModel) Insert(title, content, expires string) (int, error) {
	statement := `INSERT INTO events (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (m *EventModel) Get(id int) (*models.Event, error) {
	statement := `SELECT id, title, content, created, expires FROM events
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(statement, id)

	e := &models.Event{}

	err := row.Scan(&e.ID, &e.Title, &e.Content, &e.Created, &e.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return e, nil
}

func (m *EventModel) Latest() ([]*models.Event, error) {
	return nil, nil
}
