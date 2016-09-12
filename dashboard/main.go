package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("dashboard").Parse("<html><body>{{.}}</body></html>")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, "test")
}
