package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(Name, Description, CategoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(
		"INSERT INTO courses (id, name, description, category_id) values ($1, $2, $3, $4)", id, Name, Description, CategoryID)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        Name,
		Description: Description,
		CategoryID:  CategoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("select id, name, description, category_id from courses")

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  categoryID,
		})
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("select id, name, description, category_id from courses where category_id = $1", categoryID)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  categoryID,
		})
	}

	return courses, nil
}
