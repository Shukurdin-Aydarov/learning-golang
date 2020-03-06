package data

import (
	"log"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (user *User) CreateSession() (Session, error) {
	query :=
		`INSERT INTO sessions (uuid, email, userId, createdAt) 
		VALUES ($1, $2, $3, $4) 
		RETURING id, uuid, email, userId, createdAt`

	session := Session{}
	statement, err := Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return session, err
	}
	defer statement.Close()
	err = statement.
		QueryRow(
			createUuid(),
			user.Email,
			user.Id,
			time.Now()).
		Scan(
			&session.Id,
			&session.Uuid,
			&session.Email,
			&session.UserId,
			&session.CreatedAt)

	return session, nil
}

func (user *User) Session() (bool, error) {
	session := Session{}
	err := Db.
		QueryRow(
			"SELECT id, uuid, email, userId, createdAt FROM sessions WHERE userId = $1",
			user.Id).
		Scan(
			&session.Id,
			&session.Uuid,
			&session.Email,
			&session.UserId,
			&session.CreatedAt)

	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func (session *Session) DeleteByUuid() error {
	query := "DELETE FROM sessions WHERE uuid = $1"
	statement, err := Db.Prepare(query)
	if err != nil {
		log.Fatal(nil)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(session.Uuid)
	return nil
}

func (session *Session) User() (User, error) {
	user := User{}
	err := Db.
		QueryRow(
			"SELECT id, uuid, name, email, createdAt FROM users WHERE id = $1",
			session.Id,
		).
		Scan(
			&user.Id,
			&user.Uuid,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		)

	return user, err
}

func SessionDeleteAll() error {
	query := "DELETE FROM sessions"
	_, err := Db.Exec(query)
	return err
}

func (user *User) Create() error {
	query := `INSERT INTO users (uuid, name, email, password, createdAt)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, uuid, createdAt`

	statement, err := Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statement.Close()
	err = statement.
		QueryRow(
			createUuid(),
			user.Name,
			user.Email,
			Encrypt(user.Email),
			time.Now(),
		).
		Scan(
			&user.Id,
			&user.Uuid,
			&user.CreatedAt,
		)

	return err
}

func (user *User) Delete() error {
	query := "DELETE FROM users WHERE id = $1"
	statement, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(user.Id)

	return err
}

func (user *User) Update() error {
	query := "UPDATE users SET name = $2 email = $3 WHERE id = $1"
	statement, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(user.Id, user.Name, user.Email)

	return err
}

func (user *User) DeleteAll() error {
	query := "DELETE FROM users"
	_, err := Db.Exec(query)

	return err
}

func Users() ([]User, error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, createdAt FROM users")
	var users []User
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		scanError := rows.Scan(
			&user.Id,
			&user.Uuid,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt)

		if scanError != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func UserByEmail(email string) (User, error) {
	user := User{}
	query := "SELECT id, uuid, name, email, password, createdAt FROM users WHERE email = $1"
	statement, err := Db.Prepare(query)
	if err != nil {
		return user, err
	}
	defer statement.Close()
	err = statement.QueryRow(user.Email).
		Scan(
			&user.Id,
			&user.Uuid,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)
	return user, err
}

func UserByUuid(uuid string) (User, error) {
	user := User{}
	err := Db.
		QueryRow(
			"SELECT id, uuid, name, email, password, createdAt FROM users WHERE uuid = $1",
			uuid,
		).
		Scan(
			&user.Id,
			&user.Uuid,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)
	return user, err
}
