package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
		envErr := godotenv.Load()
		if envErr != nil {
			log.Fatal(envErr)
		}

		baseUrl := struct {
			LineBaseUrl string
		}{
			os.Getenv("LINE_BASE_URL"),
		}

		err := rnd.HTML(w, http.StatusOK, "index", baseUrl)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":80", nil)
}
