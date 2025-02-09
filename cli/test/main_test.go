package main_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Example configuration for DB container (using PostgreSQL)
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13", // Change version as needed
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	dbContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start DB container: %v", err)
	}

	// Terminate the container at the end of the test
	defer func() {
		if err := dbContainer.Terminate(ctx); err != nil {
			log.Printf("Failed to terminate DB container: %v", err)
		}
	}()

	// Get the host and mapped port of the container
	host, err := dbContainer.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get container host: %v", err)
	}
	mappedPort, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get mapped port: %v", err)
	}

	// Set DB connection info for the CLI application via environment variables
	dbURL := fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable", host, mappedPort.Port())

	// set DATABASE_URL environment variable
	// use this environment variable in the CLI application to connect to the DB
	os.Setenv("DATABASE_URL", dbURL)
	log.Println("DB connection info:", dbURL)

	// Run migration: create table and insert sample data
	if err := runMigration(dbURL); err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

// runMigration creates a table in the DB and inserts sample data
func runMigration(dbURL string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create table sample
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sample (
			id SERIAL PRIMARY KEY,
			value TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Insert sample data
	_, err = db.Exec(`INSERT INTO sample (value) VALUES ('expected output')`)
	return err
}

func TestCLICommand(t *testing.T) {
	// Adjust the path and arguments of the CLI binary according to your project
	cmd := exec.Command("go", "run", ".", "command", "--option", "value")
	// The environment variables set in TestMain are inherited, so additional settings can be added if needed

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute CLI command: %v, output: %s", err, string(out))
	}

	expected := "expected output"
	if !strings.Contains(string(out), expected) {
		t.Errorf("Expected output not found.\nOutput: %s\nExpected: %s", string(out), expected)
	}
}
