package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // _ black identify, ignore no use import
	"github.com/paulojr83/Go-Expert/SQLC/internal/db"
)

type CourseDDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDDB {
	return &CourseDDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("Error on rollback: %v, original error %w: ", errRb, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDDB) CreateCourseAndCategory(ctx context.Context, categoryParams CategoryParams, courseParams CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          categoryParams.ID,
			Name:        categoryParams.Name,
			Description: categoryParams.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          courseParams.ID,
			Name:        courseParams.Name,
			Description: courseParams.Description,
			CategoryID:  categoryParams.ID,
			Price:       courseParams.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {

	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s\nCourse ID: %s\nCourse Name: %s\nCourse Description: %s\nCourse Price: %f\n\n",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}
	/*categoryParams := CategoryParams{
		ID:   uuid.New().String(),
		Name: "Test 1",
		Description: sql.NullString{
			String: "Test 1",
			Valid:  true,
		},
	}
	courseParams := CourseParams{
		ID:   uuid.New().String(),
		Name: "Test 1",
		Description: sql.NullString{
			String: "Test 1",
			Valid:  true,
		},
		Price: 99.10,
	}

	courseDB := NewCourseDB(dbConn)

	err = courseDB.CreateCourseAndCategory(ctx, categoryParams, courseParams)

	if err != nil {
		panic(err)
	}*/
}
