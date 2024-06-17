package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type PostgreSql struct {
	db *sql.DB
}

func (postgreSql *PostgreSql) Connect() error {
	var (
		host     = os.Getenv("POSTGRESQL_HOST")
		port     = os.Getenv("POSTGRESQL_PORT")
		user     = os.Getenv("POSTGRESQL_USERNAME")
		password = os.Getenv("POSTGRESQL_PASSWORD")
		dbname   = os.Getenv("POSTGRESQL_DATABASE_NAME")
	)

	// Construct connection string
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a database connection
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return err
	}

	postgreSql.db = db

	fmt.Println("connected to PostgreSQL")
	return nil
}

func (postgreSql *PostgreSql) Close() error {
	return postgreSql.db.Close()
}

func (postgreSql *PostgreSql) BatchSaveFileComponentSummaries(fileComponentPayloads []FileComponentPayload) ([]FileComponentSummary, error) {
	if err := postgreSql.ensureFileComponentSummariesTableExists(); err != nil {
		return nil, err
	}

	fmt.Println(fileComponentPayloads)

	tx, err := postgreSql.db.Begin()
	if err != nil {
		return nil, err
	}

	statement, err := tx.Prepare(`
		INSERT INTO file_component_summaries(file_component_id, summary)
		VALUES ($1, $2)
		RETURNING id, file_component_id, summary
	`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var fileComponentSummaries []FileComponentSummary

	// Iterate over the fileComponentPayloads and add values to the batch
	for _, fileComponentPayload := range fileComponentPayloads {
		var fileComponentSummary FileComponentSummary

		err := statement.QueryRow(
			fileComponentPayload.FileComponentId,
			fileComponentPayload.Summary,
		).Scan(
			&fileComponentSummary.Id,
			&fileComponentSummary.FileComponentId,
			&fileComponentSummary.Summary,
		)

		if err != nil {
			tx.Rollback() // Rollback transaction on error
			return nil, err
		}

		fileComponentSummaries = append(fileComponentSummaries, fileComponentSummary)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return fileComponentSummaries, nil
}

func (postgreSql *PostgreSql) ensureFileComponentSummariesTableExists() error {
	createFileComponentsTableQuery := `
		CREATE TABLE IF NOT EXISTS file_component_summaries (
			id SERIAL PRIMARY KEY,
			file_component_id INT,
			summary VARCHAR(255)
		);
	`

	_, err := postgreSql.db.Exec(createFileComponentsTableQuery)

	return err
}
