package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

// DB set up
func setupDB() *sql.DB {
    DB_HOST := os.Getenv("DB_HOST")
    DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
    DB_USER := os.Getenv("DB_USER")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_NAME := os.Getenv("DB_NAME")
    
    dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return db
}

type Hub struct {
    id   int `json:"id"`
    name string `json:"name"`
    shortName string `json:"shortName"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Hub `json:"data"`
    Message string `json:"message"`
}

// Main function
func main() {

    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
      }

    // Init the mux router
    router := mux.NewRouter()

   // Route handles & endpoints
    // Get all hubs
    router.HandleFunc("/hubs/", GetHubs).Methods("GET")

    // serve the app
    fmt.Println("Server at 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}

// Function for handling messages
func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

// Get all hubs

// response and request handlers
func GetHubs(w http.ResponseWriter, r *http.Request) {
    db := setupDB()

    printMessage("Getting hubs...")

    // Get all hubs from hubs table
    rows, err := db.Query("SELECT * FROM hub")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var hubs []Hub

    // Foreach hub
    for rows.Next() {
        var id int
        var name string
        var short_name string

        err = rows.Scan(&id, &name, &short_name)

        // check errors
        checkErr(err)

        hubs = append(hubs, Hub{id: id, name: name, shortName: short_name})
    }

    var response = JsonResponse{Type: "success", Data: hubs}

    w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}