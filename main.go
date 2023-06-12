package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		file, err := ioutil.ReadFile("./assets/init.html")
		if err != nil {
			e := fmt.Sprintf("error reading html file: %s", err)
			log.Println(e)
			http.Error(w, e, http.StatusInternalServerError)
			return
		}

		w.Write(file)
	})

	// return 200 on healthz path
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthy!"))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Println("error starting server: ", err)
	}
}
