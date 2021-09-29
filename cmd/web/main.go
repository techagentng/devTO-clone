package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct{
	errLog *log.Logger
	infoLog *log.Logger
}


func main()  {
	addr := flag.String("addr", ":5300", "Pass network address ")
	flag.Parse()
	//Created custom error logs for standard out (terminal )
	    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
		errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)
	application := &application{
		errLog: errorLog,
		infoLog: infoLog,
	}

	//Creating a new server instance and parsing values as described in the standard library
	server := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: application.routes(), //r
	}
	//log.Printf("starting server at %s\n", *addr)
	infoLog.Printf("Starting our server on %s\n", *addr)
	if err := server.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}

}
