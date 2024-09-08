package main

import (
	"context"
	"fmt"
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
	err := prepare()
	if err != nil {
		log.Fatalf("Could not prepare: %s", err)
	}
	m.Run()
}

func prepare() error {
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

	return nil
}

type User struct {
	ID   string
	Name string
	Bio  string
}

func Test_Query(t *testing.T) {
	t.Parallel()

	sqlxDB := helperDB(t)

	_, err := sqlxDB.Exec("INSERT INTO users (id, name, bio) VALUES ('1', 'Alice', 'Hello'), ('2', 'Bob', 'World')")
	assert.NoError(t, err)

	rows, err := sqlxDB.QueryContext(context.Background(), "SELECT * FROM users")
	assert.NoError(t, err)

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

	var users []User
	err = sqlxDB.GetContext(context.Background(), &users, "SELECT * FROM users")
	assert.NoError(t, err)

	fmt.Println(users)

	// var user User
	// err = sqlxDB.GetContext(context.Background(), &user, "SELECT * FROM users")
	// assert.NoError(t, err)

	// assert.Equal(t, User{ID: "5", Name: "Alice", Bio: "Hello"}, user)
}

func helperDB(t *testing.T) *sqlx.DB {
	t.Helper()

	sqlxDB, err := sqlx.Open("txdb", uuid.NewString())
	assert.NoError(t, err)

	t.Cleanup(func() {
		sqlxDB.Close()
	})

	return sqlxDB
}
