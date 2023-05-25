package mysql

import (
	"database/sql"
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

	id, err := result, LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (m *EventModel) Get(id int) (*models.Event, error) {
	return nil, nil
}

func (m *EventModel) Latest() ([]*models.Event, error) {
	return nil, nil
}
