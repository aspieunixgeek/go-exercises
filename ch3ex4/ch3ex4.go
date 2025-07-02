package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var port = flag.String("p", "8000", "./ch3ex4 -p 9000")

func main() {
	flag.Parse()

	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	var width, height = 600, 320
	var color, bg = "grey", "black"

	w.Header().Set("Content-Type", "image/svg+xml")

	if query.Has("width") {
		var err error
		width, err = strconv.Atoi(query.Get("width"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if query.Has("height") {
		var err error
		height, err = strconv.Atoi(query.Get("height"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if query.Has("color") {
		color = query.Get("color")
	}

	// background color
	if query.Has("bg") {
		bg = query.Get("bg")
	}

	out := Gen(width, height, color, bg)
	_, err := w.Write([]byte(out))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
