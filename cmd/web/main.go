package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application stores the application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	// call this *before* you use the addr variable otherwise it will always contain the default value of ":3000".
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	svr := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on: %s", *addr)

	err := svr.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
