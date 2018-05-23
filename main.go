package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rfaulhaber/fdate"
	"github.com/rfaulhaber/gflag"
	"log"
	"net/http"
	"time"
)

type DateResponse struct {
	Raw           string `json:"raw"`
	Day           int    `json:"day"`
	Month         int    `json:"month"`
	Year          int    `json:"year"`
	DayOfYear     int    `json:"day_of_year"`
	Weekday       int    `json:"weekday"`
	WeekdayString string `json:"weekday_string"`
	YearString    string `json:"year_string"`
	MonthString   string `json:"month_string"`
}

func NewDateResponse(date fdate.Date) DateResponse {
	year, month, day := date.Date()

	return DateResponse{date.String(), day, int(month), year, date.DayOfYear(), int(date.Weekday()), date.Weekday().String(), date.RomanYear().String(), month.String()}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{message}
}

const dateFormat = "2006-01-02"

/*
 * TODO
 * - add godep
 */

func main() {
	port := gflag.String("p", "port", "server port to listen on", "8000")

	gflag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/today", GetToday).Methods("GET")
	router.HandleFunc("/date", GetDate).Methods("GET")

	log.Fatalln(http.ListenAndServe(":"+*port, router))
}

func GetToday(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(NewDateResponse(fdate.Today()))

	if err != nil {
		log.Println("error", err)
	}

	w.Write(response)
}

func GetDate(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	tz := r.URL.Query().Get("tz")

	if date == "" {
		message, _ := json.Marshal(NewErrorResponse("No date parameter found"))
		w.Write(message)
		return
	}

	if tz == "" {
		message, _ := json.Marshal(NewErrorResponse("No timezone parameter found"))
		w.Write(message)
		return
	}

	parsedTime, err := time.Parse(dateFormat, date) // does this deal with timezones?

	if err != nil {
		message, _ := json.Marshal(NewErrorResponse("Time is not yyyy-mm-dd format"))
		w.Write(message)
		return;
	}

	loc, err := time.LoadLocation(tz)

	log.Println(loc.String())
	log.Println(parsedTime.In(loc))

	if err != nil {
		message, _ := json.Marshal(NewErrorResponse("Timezone does not exist: " + tz))
		w.Write(message)
		return;
	}

	frcDate := fdate.DateFromTime(parsedTime.In(loc))

	response, _ := json.Marshal(NewDateResponse(frcDate))

	w.Write(response)
}
