package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rfaulhaber/fdate"
	"log"
	"net/http"
)

type DateResponse struct {
	Raw         string `json:"raw"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
	DayOfYear   int    `json:"day_of_year"`
	Weekday     string `json:"weekday"`
	YearString  string `json:"year_string"`
	MonthString string `json:"month_string"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/today", GetToday).Methods("GET")
	router.HandleFunc("/date", GetDate)
	router.HandleFunc("/date/{type}", GetDate)

	log.Fatalln(http.ListenAndServe(":8000", router))
}

func GetToday(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(responseFromDate(fdate.Today()))

	if err != nil {
		log.Println("error", err)
	}

	w.Write(response)
}

func GetDate(w http.ResponseWriter, r *http.Request) {
}

func responseFromDate(date fdate.Date) DateResponse {
	year, month, day := date.Date()

	return DateResponse{date.String(), day, int(month), year, date.DayOfYear(), date.Weekday().String(), date.RomanYear().String(), month.String()}
}
