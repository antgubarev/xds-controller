package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeout := 1000
		timeoutS := r.URL.Query().Get("timeout")
		if timeoutS != "" {
			timeoutInt, err := strconv.Atoi(timeoutS)
			if err != nil {
				w.Write([]byte("Invalid timeout param"))
				return
			}
			timeout = timeoutInt
		}
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		log.Printf("send seponse after %d millisesonds \n", timeout)
	})

	log.Print("Echo server listening on port 8080.\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
