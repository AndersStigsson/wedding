package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Vote struct {
	Id     int `json:"id"`
	BeerId int `json:"beerId"`
	UserId int `json:"userId"`
	Points int `json:"points"`
}

type AuthenticationError struct {
	Error bool   `json:"error"`
	Desc  string `json:"desc"`
}

type Beer struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Votes      []Vote  `json:"votes"`
	Points     *int    `json:"points"`
	Year       int     `json:"year"`
	PictureURL *string `json:"pictureURL"`
}

var Beers []Beer

var db *sql.DB

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func validateRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		}
		secret := os.Getenv("API_TOKEN")
		key := r.Header.Get("x-api-token")
		allowed := key == secret
		fmt.Printf("Expected %s, got %s", secret, key)
		if !allowed {
			var err AuthenticationError
			err.Desc = "Failed to authenticate"
			err.Error = true
			json.NewEncoder(w).Encode(err)
			return
		}
		next(w, r)

	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "x-api-token"})
	//originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":10010", handlers.CORS(headersOk, methodsOk)(myRouter)))
}

func returnAllGuests(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM guests")
	var beers []Beer
	if err != nil {
		log.Fatal("Error from DB")
		panic(err.Error())
	}

	for rows.Next() {
		var beer Beer
		if err := rows.Scan(&beer.Id, &beer.Name, &beer.Type, &beer.Year, &beer.PictureURL); err != nil {
			panic(err.Error())
		}
		beers = append(beers, beer)
	}

	json.NewEncoder(w).Encode(beers)
}

func createNewBeer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var beer Beer
	json.Unmarshal(reqBody, &beer)

	_, err := db.Exec("INSERT INTO beers (name, type, year, pictureURL) VALUES (?, ?, ?, ?)", beer.Name, beer.Type, beer.Year, beer.PictureURL)
	if err != nil {
		fmt.Println("Failed in insert")
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	enverr := godotenv.Load(".env")

	if enverr != nil {
		log.Fatalf("Error loading .env")
	}
	var username = os.Getenv("DBUSER")
	var password = os.Getenv("DBPASS")
	var host = os.Getenv("DBHOST")
	var port = os.Getenv("DBPORT")
	var dbname = os.Getenv("DBNAME")
	var databasestring = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	var err error
	db, err = sql.Open("mysql", databasestring)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handleRequests()

}
