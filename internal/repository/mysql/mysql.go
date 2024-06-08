package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/nurhidaylma/lover-app.git/internal/repository"

	_ "github.com/go-sql-driver/mysql"
)

type DBConf struct {
	User     string
	Password string
	URL      string
	Schema   string
}

type dbRepository struct {
	db *sql.DB
}

func (cnf *DBConf) NewRepository(r *repository.Repository) error {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&readTimeout=360s", cnf.User, cnf.Password, cnf.URL, cnf.Schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		log.Fatal("error connecting to database: ", err)
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)

	r.DB = &dbRepository{
		db: db,
	}

	return nil
}
