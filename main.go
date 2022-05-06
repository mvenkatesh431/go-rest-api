package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	/*
		HealthCheck endpoint will return the Health Status.
		By default, It provides the status as healthy. (If container running properly)
	*/
	log.Println("Entering /health-check endpoint")
	json.NewEncoder(w).Encode(map[string]bool{"healthy": true})
}

func Info(w http.ResponseWriter, r *http.Request) {
	/*
		/info endpoint will return request info.
	*/

	log.Println("Entering /info endpoint, returning info")
	remoteIP := r.RemoteAddr //requesterIP
	rMethod := r.Method
	rURI := r.RequestURI

	infoResp := make(map[string]string)
	infoResp["Version"] = os.Getenv("VERSION")
	infoResp["RemoteIP"] = remoteIP
	infoResp["Method"] = rMethod
	infoResp["Endpoint"] = rURI
	infoResp["Host"] = r.Host

	// Or you can also make the map like this.
	// infoResp := map[string]string{"Version": os.Getenv("VERSION"), "RemoteIP": remoteIP, "Method": rMethod, "Endpoint": rURI}

	json.NewEncoder(w).Encode(infoResp)
}

func Env(w http.ResponseWriter, r *http.Request) {
	/*
		/env endpoint will dump the environment variables.
	*/

	log.Println("Entering /env endpoint, Returning complete Env variables")

	infoResp := make(map[string]string)

	infoResp["Version"] = os.Getenv("VERSION")
	// dump the enviroment variables.
	// os.Environ() returns the slice of strings, So join them using strings.Join to make a string.
	infoResp["env"] = strings.Join(os.Environ(), ", ")

	json.NewEncoder(w).Encode(infoResp)
}

func main() {

	// Set the version of the API and Port
	os.Setenv("VERSION", "1.0")
	os.Setenv("PORT", "10000")

	router := mux.NewRouter()

	// router.HandleFunc("/", HomeRouter).Methods("GET")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/info", Info).Methods("GET")
	router.HandleFunc("/env", Env).Methods("GET")
	http.Handle("/", router)

	addr := "127.0.0.1:" + os.Getenv("PORT")

	log.Printf("Simple API Server(%s) running on %s", os.Getenv("VERSION"), addr)
	srv := &http.Server{
		Handler: router,
		Addr:    addr, // localhost
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
