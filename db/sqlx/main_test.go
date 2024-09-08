package main

import (
	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/blck-snwmn/example-go/db/sqlx/db"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestMain(m *testing.M) {
	err := testMain(m)
	if err != nil {
		log.Fatalf("Could not prepare: %s", err)
	}
}

func testMain(m *testing.M) error {
	ctx := context.Background()

	mysqlContainer, err := mysql.Run(
		ctx,
		"mysql:5.7",
		testcontainers.WithEnv(map[string]string{
			"MYSQL_ROOT": "secret",
		}),
	)
	if err != nil {
		return err
	}
	defer mysqlContainer.Terminate(ctx)

	connStr, err := mysqlContainer.ConnectionString(ctx)
	if err != nil {
		return err
	}
	sqlxDB, err := sqlx.Open("mysql", connStr)
	if err != nil {
		return err
	}

	if _, err := sqlxDB.Exec(db.Schema); err != nil {
		return err
	}
	txdb.Register("txdb", "mysql", connStr)

	m.Run()

	return nil
}

type User struct {
	ID   string
	Bio  string
	Name string
}

func Test_Query(t *testing.T) {
	t.Parallel()

	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec("INSERT INTO users (id, name, bio) VALUES ('1', 'Alice', 'Hello'), ('2', 'Bob', 'World')")
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

	assert.Equal(t, 2, len(users))
}

func Test_Queryx(t *testing.T) {
	t.Parallel()

	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec("INSERT INTO users (id, name, bio) VALUES ('3', 'Alice', 'Hello'), ('4', 'Bob', 'World')")
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

	assert.Equal(t, 2, len(users))
}

func Test_Get(t *testing.T) {
	t.Parallel()

	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec("INSERT INTO users (id, name, bio) VALUES ('5', 'Alice', 'Hello'), ('6', 'Bob', 'World')")
	assert.NoError(t, err)

	var user User
	err = sqlxDB.GetContext(context.Background(), &user, "SELECT * FROM users")
	assert.NoError(t, err)

	assert.Equal(t, User{ID: "5", Name: "Alice", Bio: "Hello"}, user)
}

func Test_Select(t *testing.T) {
	t.Parallel()

	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec("INSERT INTO users (id, name, bio) VALUES ('7', 'Alice', 'Hello'), ('8', 'Bob', 'World')")
	assert.NoError(t, err)

	users := []User{}
	err = sqlxDB.SelectContext(context.Background(), &users, "SELECT * FROM users")
	assert.NoError(t, err)

	assert.Equal(t, []User{
		{ID: "7", Name: "Alice", Bio: "Hello"},
		{ID: "8", Name: "Bob", Bio: "World"},
	}, users)
}

func Test_Transaction(t *testing.T) {
	t.Parallel()

	sqlxDB := helperDB(t)

	{
		tx, err := sqlxDB.Beginx()
		assert.NoError(t, err)

		_, err = tx.Exec("INSERT INTO users (id, name, bio) VALUES ('9', 'Alice', 'Hello'), ('10', 'Bob', 'World')")
		assert.NoError(t, err)

		tx.Rollback()

		var users []User
		err = sqlxDB.SelectContext(context.Background(), &users, "SELECT * FROM users")
		assert.NoError(t, err)

		assert.Equal(t, 0, len(users))
	}
	{
		tx, err := sqlxDB.Beginx()
		assert.NoError(t, err)

		_, err = tx.Exec("INSERT INTO users (id, name, bio) VALUES ('9', 'Alice', 'Hello'), ('10', 'Bob', 'World')")
		assert.NoError(t, err)

		err = tx.Commit()
		assert.NoError(t, err)

		var users []User
		err = sqlxDB.SelectContext(context.Background(), &users, "SELECT * FROM users")
		assert.NoError(t, err)

		assert.Equal(t, 2, len(users))
	}
}

func Test_In(t *testing.T) {
	t.Parallel()

	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec(`
		INSERT INTO users (id, name, bio) VALUES
		('11', 'Alice', 'Hello'),
		('12', 'Bob', 'World'),
		('13', 'Charlie', 'Hello'),
		('14', 'David', 'World')
	`)
	assert.NoError(t, err)

	query, args, err := sqlx.In("SELECT * FROM users WHERE name IN (?)", []string{"Alice", "Bob"})
	assert.NoError(t, err)

	var users []User
	err = sqlxDB.SelectContext(context.Background(), &users, sqlxDB.Rebind(query), args...)
	assert.NoError(t, err)

	assert.Equal(t, []User{
		{ID: "11", Name: "Alice", Bio: "Hello"},
		{ID: "12", Name: "Bob", Bio: "World"},
	}, users)
}

func helperDB(t *testing.T) *sqlx.DB {
	t.Helper()

	sqlxDB := sqlx.MustOpen("txdb", uuid.NewString())

	t.Cleanup(func() {
		sqlxDB.Close()
	})

	return sqlxDB
}
