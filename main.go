package main

import (
	"github.com/Molsbee/crypt/gui"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

func main() {
	app.Route("/", &gui.Hello{})
	app.Route("/hello", &gui.Hello{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "A Hello World! example",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
