package mysql

import (
	"database/sql"
	"github.com/tamudashe/un/pkg/repository"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert will insert a new snippet into the database.
func (r *SnippetModel) Insert(title, content, expires string) (int, error) {
	statement := "INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"

	result, err := r.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get will get a specific snippet from the database based on id
func (r *SnippetModel) GetSnippetByID(id int) (*repository.Snippet, error) {
	// Initialize a pointer to a new zeroed Snippet struct.
	s := &repository.Snippet{}
	statement := "SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?"

	err := r.DB.QueryRow(statement, id).
		Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, repository.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// GetLatestSnippets gets the 10 most recent snippets
func (r *SnippetModel) GetLatestSnippets() ([]*repository.Snippet, error) {
	statement := "SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10"

	rows, err := r.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*repository.Snippet{}
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		s := &repository.Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
