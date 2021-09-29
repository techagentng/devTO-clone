package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)
//This method uses the application struct in ( line10 main) to handle custom errors
func (app *application) serverError (w http.ResponseWriter, err error){
	//catch the long error printed as error
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//Attach this to our custom error
	app.errLog.Println(trace)
	http.Error(w,  http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
//Helper function to wrap client error
func (app *application) clientError (w http.ResponseWriter, status int){
	http.Error(w, http.StatusText(status), status)
}
//Helper function to wrap client error status code
func (app *application) notFound (w http.ResponseWriter){
	app.clientError(w, http.StatusNotFound)
}