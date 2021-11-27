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

type AuthenticationError struct {
	Error bool   `json:"error"`
	Desc  string `json:"desc"`
}

type Guest struct {
	Id      int     `json:"id"`
	Email   string  `json:"email"`
	Amount  int     `json:"amount"`
	Parking bool    `json:"parking"`
	Speech  bool    `json:"speech"`
	Extras  []Extra `json:"guests"`
}

type Extra struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	FoodPreference string `json:"foodpreference"`
}

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
	myRouter.HandleFunc("/add", validateRequest(addNewGuest)).Methods("POST")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "x-api-token"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":10010", handlers.CORS(headersOk, methodsOk)(myRouter)))
}

func returnAllGuests(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM guests")
	var guests []Guest
	if err != nil {
		log.Fatal("Error from DB")
		panic(err.Error())
	}

	for rows.Next() {
		var guest Guest
		if err := rows.Scan(&guest.Id, &guest.Email, &guest.Speech, &guest.Parking, &guest.Amount); err != nil {
			panic(err.Error())
		}
		guests = append(guests, guest)
	}

	json.NewEncoder(w).Encode(guests)
}

func addNewGuest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var guest Guest
	var id int
	json.Unmarshal(reqBody, &guest)

	err := db.QueryRow("INSERT INTO guests (email, speech, parking, amount) VALUES (?, ?, ?, ?) RETURNING id", guest.Email, guest.Speech, guest.Parking, guest.Amount).Scan(&id)
	if err != nil {
		fmt.Println("Failed in insert")
		panic(err.Error())
	}

	for _, element := range guest.Extras {
		_, err := db.Exec("INSERT INTO extras (guest_id, name, food_preferences) VALUES (?, ?, ?)", id, element.Name, element.FoodPreference)
		if err != nil {
			fmt.Printf("Failed to add extra %s, with fp %s on guest %d \n", element.Name, element.FoodPreference, id)
		}
	}

	json.NewEncoder(w).Encode("Success")
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	enverr := godotenv.Load(".env")

	if enverr != nil {
		log.Fatalf("Error loading .env")
	}
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")
	databasestring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	var err error
	db, err = sql.Open("mysql", databasestring)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handleRequests()
}
