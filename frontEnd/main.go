package main

import (
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	rnd = renderer.New(
		renderer.Options{
			ParseGlobPattern: "html/*.html",
		},
	)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		baseUrl := struct {
			LineBaseUrl string
		}{
			"test",
		}

		err := rnd.HTML(w, http.StatusOK, "index", baseUrl)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":80", nil)
}
