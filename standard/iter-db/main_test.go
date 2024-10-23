package db

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	err := testMain(m)
	if err != nil {
		log.Fatalf("Could not prepare: %s", err)
	}
}

var connStr string

func testMain(m *testing.M) error {
	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return err
	}
	defer postgresContainer.Terminate(ctx)

	connStr, err = postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return err
	}
	sqlxdb := sqlx.MustOpen("pgx", connStr)
	defer sqlxdb.Close()

	if _, err := sqlxdb.Exec(Schema); err != nil {
		return err
	}

	users := []User{
		{ID: "1", Name: "Alice", Bio: "Lorem ipsum"},
	}
	for i := 0; i < 100; i++ {
		users = append(users, User{
			ID:   uuid.NewString(),
			Name: uuid.NewString(),
			Bio:  uuid.NewString(),
		})
	}

	_, err = sqlxdb.NamedExec(`INSERT INTO users (id, name, bio) VALUES (:id, :name, :bio)`, users)
	if err != nil {
		return err
	}

	m.Run()

	return nil
}

type User struct {
	ID   string
	Bio  string
	Name string
}

func TestSimpleGet(t *testing.T) {
	ctx := context.Background()

	sqlxdb := sqlx.MustOpen("pgx", connStr)
	defer sqlxdb.Close()

	err := sqlxdb.PingContext(ctx)
	assert.NoError(t, err)

	var users []User
	err = sqlxdb.Select(&users, "SELECT * FROM users")
	assert.NoError(t, err)

	assert.Len(t, users, 101)
}
