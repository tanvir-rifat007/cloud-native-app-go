// Package main is the entry point to the server. It reads configuration, sets up logging and error handling,
// handles signals from the OS, and starts and stops the server.
package main

import (
	"canvas/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	os.Exit(start())
}

var release string


func start()int{
	logEnv:= getStringOrDefault("LOG_ENV", "development")
	log,err:= createLogger(logEnv)

	if err != nil {
		fmt.Println("Error creating logger:", err)
		return 1
	}

	log = log.With(zap.String("release", release))

	defer func(){
		_ = log.Sync()

	}()

	host:= getStringOrDefault("HOST", "localhost")
	port:= getIntOrDefault("PORT", 8080)

	s:= server.New(server.Options{
		Host: host,
		Port: port,
		Log : log,
	})

	

	// catch the os signals
	// sigterm and sigint

  ctx,stop:= signal.NotifyContext(context.Background(),syscall.SIGINT,syscall.SIGTERM)

	defer stop()
	
	eg,ctx:=errgroup.WithContext(ctx)

	eg.Go(func() error {

		if err:= s.Start(); err!=nil{
			 log.Error("could not start server", zap.Error(err))
			 return err
		}

		return nil
		

	})

	<-ctx.Done()

	eg.Go(func()error{
		if err:= s.Stop(); err!=nil{
			log.Error("could not stop server", zap.Error(err))
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 1
	}




  return 0;
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


func getStringOrDefault(name,defaultV string) string{
	v,ok:=os.LookupEnv(name)

	if !ok{
		return defaultV
	}
	return v
}


func getIntOrDefault(name string,defaultV int) int{
	v,ok:=os.LookupEnv(name)

	if !ok{
		return defaultV
	}

	intV,err:=strconv.Atoi(v)

	if err!=nil{
		return defaultV
	}
	return intV
}
