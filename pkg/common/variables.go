package common

import (
	"database/sql"
	"log"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Debug    bool
	WebURL   string
	// ApiURL   string
	// Payment  *payevents.PayEvents
}

var Db *sql.DB
var App Application
