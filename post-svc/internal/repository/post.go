package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// QueryRowSQL is a wrapper function that logs the SQL command before executing it.
func QueryRowSQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (*sql.Row, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided arguments
	row := stmt.QueryRowContext(ctx, args...)

	return row, nil
}

// QuerySQL is a wrapper function that logs the SQL command before executing it.
func QuerySQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (*sql.Rows, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)
	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided arguments
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Printf("Error executing SQL command: %s", err.Error())
		return nil, err
	}

	return rows, nil
}

// ExecSQL is a wrapper function that logs the SQL command before executing it.
func ExecSQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (sql.Result, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)
	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL command with the provided arguments
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Printf("Error executing SQL command: %s", err.Error())
		return nil, err
	}

	return result, nil
}

type PostInputRepo struct {
	ID               int
	Title            string
	Content          string
	ClassroomID      int
	ReportingStageID int
	AuthorID         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// CreatePost creates a new post in db given by post model
func (r *PostRepo) CreatePost(ctx context.Context, p PostInputRepo) (PostOutputRepo, error) {
	// check post exists
	// isExists, err := r.IsPostExists(ctx, p.Title, p.ClassroomID)
	// if err != nil {
	// 	return PostOutputRepo{}, err
	// }

	// if isExists {
	// 	return PostOutputRepo{}, ErrPostExisted
	// }

	row, err := QueryRowSQL(ctx, r.Database, "CreatePost", "INSERT INTO posts (title, content, classroom_id, reporting_stage_id, author_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, content, classroom_id, reporting_stage_id, author_id, created_at, updated_at", p.Title, p.Content, p.ClassroomID, p.ReportingStageID, p.AuthorID)
	if err != nil {
		return PostOutputRepo{}, err
	}

	post := PostOutputRepo{}
	if err = row.Scan(&post.ID, &post.Title, &post.Content, &post.ClassroomID, &post.ReportingStageID, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return PostOutputRepo{}, ErrPostNotFound
		}
		return PostOutputRepo{}, err
	}

	return post, nil
}

type PostOutputRepo struct {
	ID               int
	Title            string
	Content          string
	ClassroomID      int
	ReportingStageID int
	AuthorID         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// GetPost returns a post in db given by id
func (r *PostRepo) GetPost(ctx context.Context, id int) (PostOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetPost", "SELECT id, title, content, classroom_id, reporting_stage_id, author_id, created_at, updated_at FROM posts WHERE id=$1", id)
	if err != nil {
		return PostOutputRepo{}, err
	}

	post := PostOutputRepo{}
	if err = row.Scan(&post.ID, &post.Title, &post.Content, &post.ClassroomID, &post.ReportingStageID, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return PostOutputRepo{}, ErrPostNotFound
		}
		return PostOutputRepo{}, err
	}

	return post, nil
}

// CheckPostExists checks whether the specified post exists by title (true == exist)
func (r *PostRepo) IsPostExists(ctx context.Context, title string, classroomID int) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsPostExists", "SELECT EXISTS(SELECT 1 FROM posts WHERE title LIKE '%' || $1 || '%' AND classroom_id=$2)", title, classroomID)
	if err != nil {
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

// UpdatePost updates the specified post by id
func (r *PostRepo) UpdatePost(ctx context.Context, id int, post PostInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdatePost", "UPDATE posts SET title=$2, content=$3, classroom_id=$4, reporting_stage_id=$5, author_id=$6, updated_at=$7 WHERE id=$1", id, post.Title, post.Content, post.ClassroomID, post.ReportingStageID, post.AuthorID, time.Now())
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrPostNotFound
	}

	return nil
}

// DeletePost deletes a post in db given by id
func (r *PostRepo) DeletePost(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeletePost", "DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrPostNotFound
	}

	return nil
}

type PostFilterRepo struct {
	Limit       int
	Page        int
	TitleSearch string
	SortColumn  string
	SortOrder   string
}

// GetPost returns a list of posts in db with filter
func (r *PostRepo) GetPosts(ctx context.Context, filter PostFilterRepo) ([]PostOutputRepo, int, error) {
	query := []string{"SELECT id, title, content, classroom_id, reporting_stage_id, author_id, created_at, updated_at FROM posts"}

	if filter.TitleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE UPPER(title) LIKE UPPER('%s')", "%"+filter.TitleSearch+"%"))
	}

	query = append(query, fmt.Sprintf("ORDER BY %s %s", filter.SortColumn, filter.SortOrder),
		fmt.Sprintf("LIMIT %d OFFSET %d", filter.Limit, (filter.Page-1)*filter.Limit))

	rows, err := QuerySQL(ctx, r.Database, "GetPosts", strings.Join(query, " "))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the posts slice
	var posts []PostOutputRepo
	for rows.Next() {
		post := PostOutputRepo{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.ClassroomID,
			&post.ReportingStageID,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCount(ctx, filter.TitleSearch)
	if err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}

// GetAllPostsOfClassroom returns all posts of the specified classroom given by classroom id
func (r *PostRepo) GetAllPostsOfClassroom(ctx context.Context, filter PostFilterRepo, classromID int) ([]PostOutputRepo, int, error) {
	query := []string{"SELECT id, title, content, classroom_id, reporting_stage_id, author_id, created_at, updated_at FROM posts"}

	if filter.TitleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d AND UPPER(title) LIKE UPPER('%s')", classromID, "%"+filter.TitleSearch+"%"))
	} else {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d", classromID))
	}

	query = append(query, fmt.Sprintf("ORDER BY %s %s", filter.SortColumn, filter.SortOrder),
		fmt.Sprintf("LIMIT %d OFFSET %d", filter.Limit, (filter.Page-1)*filter.Limit))

	rows, err := QuerySQL(ctx, r.Database, "GetAllPostsOfClassroom", strings.Join(query, " "))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the posts slice
	var posts []PostOutputRepo
	for rows.Next() {
		post := PostOutputRepo{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.ClassroomID,
			&post.ReportingStageID,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCountInClassroom(ctx, filter.TitleSearch, classromID)
	if err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}

// GetAllPostsInReportingStage returns all posts of the specified reporting stage given by reporting stage id
func (r *PostRepo) GetAllPostsInReportingStage(ctx context.Context, reportingStageID, classroomID int) ([]PostOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAllPostsInReportingStage", fmt.Sprintf("SELECT id, title, content, classroom_id, reporting_stage_id, author_id, created_at, updated_at FROM posts WHERE reporting_stage_id = %d AND classroom_id = %d", reportingStageID, classroomID))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the posts slice
	var posts []PostOutputRepo
	for rows.Next() {
		post := PostOutputRepo{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.ClassroomID,
			&post.ReportingStageID,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCountInReportingStage(ctx, reportingStageID, classroomID)
	if err != nil {
		return nil, 0, err
	}

	return posts, count, nil
}

func (r *PostRepo) getCount(ctx context.Context, titleSearch string) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM posts"}
	if titleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE UPPER(title) LIKE UPPER('%s')", "%"+titleSearch+"%"))
	}

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", strings.Join(query, " "))
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostRepo) getCountInClassroom(ctx context.Context, titleSearch string, classroomID int) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM posts"}
	if titleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d AND UPPER(title) LIKE UPPER('%s')", classroomID, "%"+titleSearch+"%"))
	} else {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d", classroomID))
	}

	rows, err := QueryRowSQL(ctx, r.Database, "getCountIntClassroom", strings.Join(query, " "))
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostRepo) getCountInReportingStage(ctx context.Context, reportingStageID, classroomID int) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCountIntClassroom", fmt.Sprintf("SELECT COUNT(*) FROM posts WHERE reporting_stage_id = %d AND classroom_id = %d", reportingStageID, classroomID))
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
