package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Checkout struct {
	gorm.Model
	ClientName string
	Amount     string
}

type CheckoutData struct {
	Name    string
	Address string
	City    string
	State   string
	Email   string
	Zip     string
	Amount  string
}

func main() {

	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/checkout", checkoutHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Service Started on Port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// db : func db
func db(checkoutData CheckoutData) {
	db, err := gorm.Open("sqlite3", "checkout.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Checkout{})

	// Create
	db.Create(&Checkout{ClientName: checkoutData.Name, Amount: checkoutData.Amount})
}

// enableCors : enable cors
func enableCors(w *http.ResponseWriter) {
	origin := os.Getenv("APP_ORIGIN_URL")
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
}

// checkoutHandler : checkout create
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	buf, bodyErr := ioutil.ReadAll(r.Body)

	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
		return
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = rdr1
	data, _ := ioutil.ReadAll(r.Body)
	//log.Println(string(data))
	var checkoutData CheckoutData
	json.Unmarshal([]byte(string(data)), &checkoutData)
	db(checkoutData)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"checkout":"ok"}`)
}

// StatusHandler : Example to make other handlers
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
