package main

import (
  "os"
  "fmt"
  "log"
  "bookbytes/books"
  "bookbytes/web"
)

func main() {
  //batching()
  filename := "./books/36-0.txt"
  book := books.BookBuilder(filename)
  books.BookPrinter(book)
  fmt.Println()
  //book.paragraph = books.GetParagraph(book)
  //fmt.Println("Index: ", book.paragraph)
  fmt.Println(books.ParagraphPrinter(book))
  fmt.Println()
  //book.paragraph = NextParagraph(book)
  //fmt.Println("Index: ", book.paragraph)
  //fmt.Println(ParagraphPrinter(book))
  //r := web.SetupRouter(book)
  //r.Run(":8080")
  web.SetupRouter()
}

// runs recursively over folder 
// mostly for testing, I think
func batching() {
  dirname := "./books"

  f, err := os.Open(dirname)
  if err != nil {
    log.Fatal(err)
  }
  files, err := f.Readdir(-1)
  f.Close()
  if err != nil {
    log.Fatal(err)
  }

  for _, file := range files {
    filename := "./books/" + file.Name()
    book := books.BookBuilder(filename)
    books.BookPrinter(book)
    fmt.Println()
  }
}
