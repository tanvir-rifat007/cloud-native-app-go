package integrationtest

import (
	"canvas/server"
	"net/http"
	"testing"
	"time"
)


func CreateServer() func(){
	s:= server.New(server.Options{
		Host: "localhost",
		Port: 8081,
	})

	// another goroutine to start the server
	go func(){
		if err:= s.Start();err!=nil{
			panic(err)
		}

	}()


	for{
		// check if server is up
		_, err:= http.Get("http://localhost:8081/")
		if err==nil{
			break;
		}
		time.Sleep(5 * time.Millisecond)
	}



	return func() {
		// shutdown the server
		if err:= s.Stop();err!=nil{
			panic(err)
		}

	}
}

// skip the test if short flag is set
func SkipIfShort(t *testing.T){
	if testing.Short(){
		t.SkipNow()
	}
}