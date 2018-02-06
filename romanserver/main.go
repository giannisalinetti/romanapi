// Example roman number server from the book
// Building RESTful Web Services with Go Author: Naren Yellavula Publisher: Packt Publishing
// Adaptment: Gianni Salinetti <gbsalinetti@extraordy.com>

package main

import (
	"fmt"
	"github.com/gbs/romanapi/romanNumerals"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const version = "v1.2"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// If request is GET with correct syntax
		if urlPathElements[1] == "roman" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "Roman numeral for %d is %q\n", number, html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else if urlPathElements[1] == "version" {
			fmt.Fprintf(w, "Roman numbers converter %s\n", version)
		} else {
			// For all other requests, tell that Client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Printing logs on standard output is ideal for a containerized environment
	log.Printf("Listening on port %s", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
