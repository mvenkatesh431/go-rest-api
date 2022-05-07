package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type JQuotes struct {
	QList []JQuote `json:"quotes"`
}

type JQuote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

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

	resp := make(map[string]string)

	resp["Version"] = os.Getenv("VERSION")
	// dump the enviroment variables.
	// os.Environ() returns the slice of strings, So join them using strings.Join to make a string.
	resp["env"] = strings.Join(os.Environ(), ", ")

	json.NewEncoder(w).Encode(resp)
}

func Quote(w http.ResponseWriter, r *http.Request) {
	/*
		/quote endpoint will return random quote.
	*/

	fileName, ok := getEnv("QUOTESFILE")
	if !ok {
		fileName = "./quotes.json"
	}

	log.Println("Entered /quote endpoint, using", fileName, "as source - Returning a Random Quote")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Not able to open %s file", fileName)
	}

	jsonQuotes, err := getJsonQuotes(file)
	if err != nil {
		log.Fatalf("Failed to decode the JSON %s", err)
	}

	// Get a random quote from "jsonQuotes"
	// Pass the time as the seed for random generator
	rand.Seed(time.Now().UnixNano())
	// Generate random index from 0 to lenght of the quote array
	rNum := rand.Intn(len(jsonQuotes.QList) - 1)
	rQuote := jsonQuotes.QList[rNum]

	// Prepare and Send the response back
	resp := make(map[string]string)
	resp["Version"] = os.Getenv("VERSION")
	resp["Quote"] = rQuote.Text + " - " + rQuote.Author
	log.Println(resp["Quote"])
	json.NewEncoder(w).Encode(resp)

}

// Parse json file and decode into the Story struct
func getJsonQuotes(file io.Reader) (JQuotes, error) {
	d := json.NewDecoder(file)
	var jsonQuotes JQuotes
	if err := d.Decode(&jsonQuotes); err != nil {
		return JQuotes{}, err
	}
	return jsonQuotes, nil
}

func getEnv(key string) (string, bool) {
	val, ok := os.LookupEnv(key)
	if !ok {
		// 'key' like "PORT" env not set return default value "10000"
		return "", false
	} else {
		return val, true
	}
}

func main() {

	// Set the version of the API and Port
	os.Setenv("VERSION", "1.0")

	router := mux.NewRouter()

	// router.HandleFunc("/", HomeRouter).Methods("GET")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/info", Info).Methods("GET")
	router.HandleFunc("/env", Env).Methods("GET")
	router.HandleFunc("/quote", Quote).Methods("GET")
	http.Handle("/", router)

	port, ok := getEnv("PORT")
	if !ok {
		port = "10000"
	}

	addr := "127.0.0.1:" + port

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
