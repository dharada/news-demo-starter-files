package main

import (
	"net/http"
	"net/http/httputil"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	requestBytes, error := httputil.DumpRequestOut(r, true)

	//if error != nil {
	//	w.Write([]byte(error.Error()))
	//}

	w.Write([]byte("<h1>Hello World!" + string(requestBytes) + r.UserAgent() + "</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}

