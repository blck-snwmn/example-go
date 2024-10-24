package db

import (
	"context"
	"iter"
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
			ID:   uuid.Must(uuid.NewV7()).String(),
			Name: uuid.Must(uuid.NewV7()).String(),
			Bio:  uuid.Must(uuid.NewV7()).String(),
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

func TestIter(t *testing.T) {
	sqlxdb := sqlx.MustOpen("pgx", connStr)
	defer sqlxdb.Close()

	var (
		dataCount int
		loopCount int
	)

	seqer := &DBSeqer{}
	for users := range seqer.Seq(sqlxdb, 10) {
		dataCount += len(users)
		loopCount++
	}
	assert.Equal(t, 101, dataCount)
	assert.Equal(t, 11, loopCount)
	assert.Equal(t, 11, seqer.LoopCouter)
	assert.NoError(t, seqer.Err)
}

func TestIter_just(t *testing.T) {
	sqlxdb := sqlx.MustOpen("pgx", connStr)
	defer sqlxdb.Close()

	var (
		dataCount int
		loopCount int
	)

	seqer := &DBSeqer{}
	for users := range seqer.Seq(sqlxdb, 101) {
		dataCount += len(users)
		loopCount++
	}
	assert.Equal(t, 101, dataCount)
	assert.Equal(t, 1, loopCount)
	assert.Equal(t, 1, seqer.LoopCouter)
	assert.NoError(t, seqer.Err)
}

type DBSeqer struct {
	Err        error
	LoopCouter int
}

func (d *DBSeqer) Seq(sqlxdb *sqlx.DB, limit int) iter.Seq[[]User] {
	return func(yield func([]User) bool) {
		offset := 0
		for {
			d.LoopCouter++

			var (
				users []User
				next  bool
			)
			users, next, d.Err = d.selectData(sqlxdb, limit, offset)
			if d.Err != nil {
				return
			}

			if !yield(users) {
				return
			}

			if !next {
				return
			}

			offset += limit
		}
	}
}

func (d *DBSeqer) selectData(sqlxdb *sqlx.DB, limit, offset int) ([]User, bool, error) {
	var users []User
	err := sqlxdb.Select(&users, "SELECT * FROM users LIMIT $1 OFFSET $2", limit+1, offset)
	if err != nil {
		return nil, false, err
	}

	if len(users) == 0 {
		return nil, false, nil
	}

	end := min(len(users), limit)
	return users[:end], len(users) > limit, nil
}
