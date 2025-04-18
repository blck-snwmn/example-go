package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// run parses command line arguments and executes the appropriate subcommand
func run(args []string) error {
	// Display usage if no subcommand is specified
	if len(args) < 1 {
		usage()
		return fmt.Errorf("subcommand is required")
	}

	switch args[0] {
	case "command":
		return runCommand(args[1:])
	default:
		usage()
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

// usage displays a simple usage message
func usage() {
	fmt.Println("Usage:")
	fmt.Println("  your_cli_binary command --option value")
}

// runCommand processes the "command" subcommand
func runCommand(args []string) error {
	// Define flags
	cmd := flag.NewFlagSet("command", flag.ExitOnError)
	option := cmd.String("option", "", "option value")
	if err := cmd.Parse(args); err != nil {
		return err
	}

	// Get DB connection info from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("environment variable DATABASE_URL is not set")
	}

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	defer db.Close() //nolint:errcheck // Closing DB connection on defer is standard practice

	// Check connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	var query string
	// If --option is specified, retrieve value from sample table
	if *option != "" {
		query = "SELECT value FROM sample ORDER BY id LIMIT 1"
	} else {
		query = "SELECT 'default output'::text"
	}

	var result string
	if err := db.QueryRow(query).Scan(&result); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	fmt.Println(result)
	return nil
}
