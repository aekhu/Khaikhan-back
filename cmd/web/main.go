package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	"git.bolor.net/orgil/go-gin-template/pkg/common"
	"git.bolor.net/orgil/go-gin-template/pkg/userman"
	"git.bolor.net/orgil/go-gin-template/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	err := os.MkdirAll("./runtime", 0755)
	if err != nil {
		panic(err)
	}
	info, err := os.Create("./runtime/info.log")
	if err != nil {
		panic(err)
	}
	errr, err := os.Create("./runtime/err.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(info, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(errr, os.Stderr)
	conf, err := util.LoadConfig("./config")
	if err != nil {
		panic(err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	common.App = common.Application{
		Debug:    conf.Debug,
		WebURL:   conf.Service.WebURL,
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		// ApiURL:   conf.Service.ApiURL,
	}
	connStr := "user=" + conf.DB.DBUser + " dbname=" + conf.DB.DBName + " password=" + conf.DB.DBPassword +
		" host=" + conf.DB.DBHost + " port=" + conf.DB.DBPort + " sslmode=" + conf.DB.DBSSLMode
	common.Db, err = openDB(connStr)
	if err != nil {
		panic(err)
	}
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "token", "content-type"},
	}))
	userman.Routers(route)

	route.Run(conf.Service.Port)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
