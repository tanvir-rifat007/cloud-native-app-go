// Package main is the entry point to the server. It reads configuration, sets up logging and error handling,
// handles signals from the OS, and starts and stops the server.
package main

import (
	"canvas/internal/data"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"

	"go.uber.org/zap"
)

var release string


type config struct{
	Port string
	Env  string
	db   struct{
		DSN string
	}
}

type application struct{
	logger *zap.Logger
	config config
	models data.Model
}

func main(){

	var cfg config

	flag.StringVar(&cfg.Port, "port","8080", "API server port")

	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|production)")

	flag.StringVar(&cfg.db.DSN, "dsn", os.Getenv("DB_URL"), "PostgreSQL DSN")

	flag.Parse()

	// create a zap logger
	log,err:= createLogger(cfg.Env)

	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}

	log = log.With(zap.String("release", release))

	defer func(){
		_ = log.Sync()

	}()


		// open a connection to the database
	db, err := openDB(cfg)

	if err != nil {
		log.Error("Error opening database connection", zap.Error(err))
		os.Exit(1)
	}

	defer db.Close()

	log.Info("Database connection established", zap.String("dsn", cfg.db.DSN))


	

	// create an application struct

	app := &application{
		logger: log,
		config: cfg,
		models: data.NewModel(db),
	}





	err = app.serve(":" + cfg.Port)

	if err != nil {
		app.logger.Error("Error starting server", zap.Error(err))
		os.Exit(1)
	}

}

func createLogger(env string) (*zap.Logger,error){
	switch env{
		case "development":
			return zap.NewDevelopment()
		case "production":
			return zap.NewProduction()
		
		default:
			return zap.NewNop(),nil
	}
}


func openDB(cfg config)(*sql.DB,error){
	db,err:= sql.Open("postgres",cfg.db.DSN)

	if err!=nil{
		return nil,err
	}

	ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err!=nil{
		return nil,err
	}

	return db,nil
}


