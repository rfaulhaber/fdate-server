package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rfaulhaber/fdate"
	"log"
	"net/http"
	"github.com/rfaulhaber/gflag"
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

/*
 * TODO
 * - accept dates, convert to FRC
 * - accept timezone, figure out current day in that timezone
 */

func main() {
	port := gflag.String("p", "port", "server port to listen on", "8000")

	gflag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/today", GetToday).Methods("GET")
	router.HandleFunc("/date", GetDate)
	router.HandleFunc("/date/{type}", GetDate)

	log.Fatalln(http.ListenAndServe(":" + *port, router))
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
