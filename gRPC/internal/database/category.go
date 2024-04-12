package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description ) values ($1, $2, $3)",
		id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) GetCategories() ([]Category, error) {
	rows, err := c.db.Query("select id, name, description from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{
			ID:          id,
			Name:        name,
			Description: description,
		})
	}
	return categories, err
}

func (c *Category) GetCategoryById(categoryId string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("select id, name, description from categories where id = $1",
		categoryId).Scan(&id, &name, &description)

	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) GetCategoryByCourseId(curseId string) (Category, error) {
	var id, name, description string

	err := c.db.QueryRow("select c.id, c.name, c.description from categories c join courses co on(c.id = co.category_id ) where co.id =  $1",
		curseId).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}
