package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/blck-snwmn/example-go/db/sqlx/db"
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
	defer postgresContainer.Terminate(ctx) //nolint:errcheck

	connStr, err := postgresContainer.ConnectionString(ctx)
	if err != nil {
		return err
	}
	sqlxDB, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return err
	}

	if _, err := sqlxDB.Exec(db.Schema); err != nil {
		return err
	}
	txdb.Register("txdb", "pgx", connStr)

	// txdb does not exist in `sqlx.defaultBinds`, so bind
	sqlx.BindDriver("txdb", sqlx.DOLLAR)

	m.Run()

	return nil
}

type User struct {
	ID   string
	Bio  string
	Name string
}

const insertSQL = `
	INSERT INTO users (id, name, bio) 
	VALUES
		 ('1', 'Alice', 'Hello'), 
		 ('2', 'Bob', 'World'),
		 ('3', 'Charlie', 'Goodbye')
`

func Test_Query(t *testing.T) {
	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec(insertSQL)
	assert.NoError(t, err)

	rows, err := sqlxDB.QueryContext(context.Background(), "SELECT * FROM users")
	assert.NoError(t, err)

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Bio)
		assert.NoError(t, err)

		users = append(users, user)
	}
	err = rows.Err()
	assert.NoError(t, err)

	assert.Equal(t, 3, len(users))
}

func Test_Queryx(t *testing.T) {
	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec(insertSQL)
	assert.NoError(t, err)

	rows, err := sqlxDB.QueryxContext(context.Background(), "SELECT * FROM users")
	assert.NoError(t, err)

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.StructScan(&user)
		assert.NoError(t, err)

		users = append(users, user)
	}
	err = rows.Err()
	assert.NoError(t, err)

	assert.Equal(t, 3, len(users))
}

func Test_Get(t *testing.T) {
	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec(insertSQL)
	assert.NoError(t, err)

	var user User
	err = sqlxDB.GetContext(
		context.Background(),
		&user,
		"SELECT * FROM users  where id = $1",
		"1",
	)
	assert.NoError(t, err)

	assert.Equal(t,
		User{ID: "1", Name: "Alice", Bio: "Hello"},
		user,
	)
}

func Test_Select(t *testing.T) {
	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec(insertSQL)
	assert.NoError(t, err)

	users := []User{}
	err = sqlxDB.SelectContext(
		context.Background(),
		&users,
		"SELECT * FROM users order by id",
	)
	assert.NoError(t, err)

	assert.Equal(t, []User{
		{ID: "1", Name: "Alice", Bio: "Hello"},
		{ID: "2", Name: "Bob", Bio: "World"},
		{ID: "3", Name: "Charlie", Bio: "Goodbye"},
	}, users)
}

func Test_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		sqlxDB := helperDB(t)
		tx, err := sqlxDB.Beginx()
		assert.NoError(t, err)

		_, err = tx.Exec(insertSQL)
		assert.NoError(t, err)

		tx.Rollback() //nolint:errcheck

		var users []User
		err = sqlxDB.SelectContext(context.Background(), &users, "SELECT * FROM users")
		assert.NoError(t, err)

		assert.Equal(t, 0, len(users))
	})

	t.Run("commit", func(t *testing.T) {
		sqlxDB := helperDB(t)
		tx, err := sqlxDB.Beginx()
		assert.NoError(t, err)

		_, err = tx.Exec(insertSQL)
		assert.NoError(t, err)

		err = tx.Commit()
		assert.NoError(t, err)

		var users []User
		err = sqlxDB.SelectContext(context.Background(), &users, "SELECT * FROM users")
		assert.NoError(t, err)

		assert.Equal(t, 3, len(users))
	})
}

func Test_In(t *testing.T) {
	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec(insertSQL)
	assert.NoError(t, err)

	query, args, err := sqlx.In("SELECT * FROM users WHERE name IN (?)", []string{"Alice", "Bob"})
	assert.NoError(t, err)

	var users []User
	err = sqlxDB.SelectContext(context.Background(), &users, sqlxDB.Rebind(query), args...)
	assert.NoError(t, err)

	assert.Equal(t, []User{
		{ID: "1", Name: "Alice", Bio: "Hello"},
		{ID: "2", Name: "Bob", Bio: "World"},
	}, users)
}

func helperDB(t *testing.T) *sqlx.DB {
	t.Helper()
	uid := uuid.NewString()
	sqlxDB := sqlx.MustOpen("txdb", uid)

	t.Cleanup(func() {
		sqlxDB.Close()
	})

	return sqlxDB
}
