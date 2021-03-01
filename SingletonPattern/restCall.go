package main

import (
	"SingletonPattern/Logger"
	"fmt"
	"net/http"
)

func main() {
	log := Logger.GetLoggerInstance()
	fmt.Println(log)
	log1 := Logger.GetLoggerInstance()
	fmt.Println(log1)
	log.Println("Started Hydra Application")

	log.Println("Response Written")
	http.HandleFunc("/sayhello", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Hydra", nil)
}
