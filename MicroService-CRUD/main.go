package main

import (
	"MICROSERVICE-CRUD/handler"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	thisLogger := log.New(os.Stdout, "product-api", log.LstdFlags)
	thisProductHandler := handler.NewProducts(thisLogger)
	thisServeMux := http.NewServeMux()
	thisServeMux.Handle("/", thisProductHandler)
	//create server
	thisServer := &http.Server{
		Addr:         ":8076",
		Handler:      thisServeMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
// go routine to start server in unclocing way 
	go func(){
		errWhileStartingServer := thisServer.ListenAndServe()
		if errWhileStartingServer != nil{
			thisLogger.Fatal(errWhileStartingServer)
		}
	}()
	signChan := make(chan os.Signal)
	//lisyen to kill interupt signals 
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	thisSignalChannel := <- signChan

	thisLogger.Println("Received Terminate, Graceful Shutdown", thisSignalChannel)
	timeOutContext, canFunct := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer canFunct()
	thisServer.Shutdown(timeOutContext)
}