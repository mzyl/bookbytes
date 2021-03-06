package web

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
  "bookbytes/books"
)

// Load the index.html template.
var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))

func SetupRouter() {
  // Serve / with the index.html file.
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  // Serve /callme with a text response.
  http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
    paragraph := books.NewParagraph()
    fmt.Fprintln(w, paragraph)
  })

  // Start the server at http://localhost:9000
  log.Fatal(http.ListenAndServe(":9000", nil))
}
