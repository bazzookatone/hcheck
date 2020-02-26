package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type HealthState struct {
	State         int
	ErrorMessages []string
}

func main() {

	http.HandleFunc("/hcheck", handleRequest)

	http.ListenAndServe(":9200", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	outputState := &HealthState{}
	checkMysql(outputState)

	if len(outputState.ErrorMessages) > 0 {
		outputState.State = 503
	} else {
		outputState.State = 200
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(outputState.State)
}


func checkMysql(outputState *HealthState) {
	db, err := sql.Open("mysql", "mysqlchkusr:ewrq89hCVBzX@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		outputState.ErrorMessages = append(outputState.ErrorMessages, fmt.Sprintf("MYSQL: %s", err.Error()))
	} else {
		defer db.Close()
	}
	err = db.Ping()
	if err != nil {
		outputState.ErrorMessages = append(outputState.ErrorMessages, fmt.Sprintf("MYSQL: %s", err.Error()))
	}
}
