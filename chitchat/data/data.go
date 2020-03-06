package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
)

// Db is initailized database
var Db *sql.DB

func init() {
	db, err := sql.Open("postgres", "dbname=chitchat ssmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	Db = db
}

func createUuid() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F

	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// Encrypt text with SHA-1
func Encrypt(text string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
}
