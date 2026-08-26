package main

import (
	"database/sql"
	"net/http"

	"github.com/heesuk/wedding-invitation-server/env"
	"github.com/heesuk/wedding-invitation-server/httphandler"
	"github.com/heesuk/wedding-invitation-server/sqldb"
	"github.com/rs/cors"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./sql.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/guestbook", httphandler.GuestbookHandler)
	mux.HandleFunc("/api/attendance", httphandler.AttendanceHandler)

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{env.AllowOrigin},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials: true,
	})

	handler := corHandler.Handler(mux)

	http.ListenAndServe(":8080", handler)
}
