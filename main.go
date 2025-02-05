package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

// Struktura podataka
type Message struct {
	Text string
}

// Render funkcija
func renderTemplate(c echo.Context, templateName string, data interface{}) error {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		return err
	}
	return tmpl.Execute(c.Response().Writer, data)
}

func main() {
	e := echo.New()

	// Ruta za prikaz početne strane
	e.GET("/", func(c echo.Context) error {
		return renderTemplate(c, "index.html", nil)
	})

	// Ruta za ažuriranje sadržaja koristeći HTMX
	e.GET("/message", func(c echo.Context) error {
		data := Message{Text: "Pozdrav iz Go i HTMX!"}
		return renderTemplate(c, "message.html", data)
	})

	// Pokretanje servera
	e.Logger.Fatal(e.Start(":8081"))
}
