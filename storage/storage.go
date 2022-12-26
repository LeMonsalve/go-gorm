package storage

import (
	"fmt"
	"log"
	"sync"

	// ...
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB // a pointer to gorm.DB
	once sync.Once
)

// New create the connection with db
func New() {
	newPostgresDB()
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", "postgres://lemonsalve:1040033687Fam@localhost:5432/GoDB?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB { // returns a pointer to gorm.DB
	return db
}
