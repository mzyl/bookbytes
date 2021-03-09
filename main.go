package main

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "strings"
)

func main() {
  //batching()
  filename := "./books/11-0.txt"
  book := BookBuilder(filename)
  BookPrinter(book)
  fmt.Println()
  fmt.Printf("%v\n", GetParagraph(book.fullText))
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
    book := BookBuilder(filename)
    BookPrinter(book)
    fmt.Println()
  }
}

type Book struct {
  filename string
  title string
  author string
  date string // release dates are really wonky.. might not include
  fullText []string
}

func BookBuilder(filename string) Book {
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  title := GetTitle(scanner)
  author := GetAuthor(scanner)
  date := GetDate(scanner)
  fullText := Strip(GetText(scanner))
  return Book{filename, title, author, date, fullText}

}

func BookPrinter(book Book) {
  fmt.Println("Filename: ", book.filename)
  fmt.Println("Title: \t\t", book.title)
  fmt.Println("Author: \t", book.author)
  fmt.Println("Release Date: \t", book.date)
  //fmt.Println("Full Text: \n", book.fullText)
}

// retrieve book title from file
func GetTitle(scanner *bufio.Scanner) (title string) {
  for scanner.Scan() {
    if strings.Contains(scanner.Text(), "Title") {
      line := strings.SplitAfter(scanner.Text(), ":")
      title = strings.TrimSpace(strings.Join(line[1:], " "))
      break;
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return 
}

// retrieve book author from file
func GetAuthor(scanner *bufio.Scanner) (author string) {
  for scanner.Scan() {
    if strings.Contains(scanner.Text(), "Author") {
      line := strings.SplitAfter(scanner.Text(), ":")
      author = strings.TrimSpace(strings.Join(line[1:], " "))
      break;
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return 
}

// retrieve book release date from file
// dates are actually weird, though, so this will probably not be used
func GetDate(scanner *bufio.Scanner) (date string) {
  for scanner.Scan() {
    if strings.Contains(scanner.Text(), "Release Date") {
      line := Between(scanner.Text(), ":", "[")
      date = strings.TrimSpace(line)
      break;
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return 
}

// helps retrieve release date, which is between a colon and a book number encased in brackets
// since the dates are weird, this will probably be removed later
func Between(line string, a string, b string) (date string) {
  first := strings.Index(line, a)
  if first == -1 {
    return ""
  }

  last := strings.Index(line, b)
  if last == -1 {
    return ""
  }

  firstAdjusted := first + len(a)
  if firstAdjusted >= last {
    return ""
  }
  date = line[firstAdjusted:last]
  return
}

// TODO:
// this needs to section by newlines or something
// currently splits line by line which does not work for returning paragraphs later
// return full text of file
func GetText(scanner *bufio.Scanner) (text []string) {
  for scanner.Scan() {
    text = append(text, scanner.Text())
  }
  return
}

// strip licensing info, index, etc. from file
func Strip(text []string) (stripped []string) {
  start := 0
  end := 0

  for i, line := range text {
    if strings.Contains(line, "***") {
      start = i
      break
    }
  }

  stripped = text[start+1:]

  // wish I could go from the bottom of the file easily, but it doesn't look like I can
  // Go does this quickly, though
  for i, line := range stripped {
    if strings.Contains(line, "***") {
      end = i
      break
    }
  }

  stripped = stripped[:end]
  return
}

// TODO:
// randomly select a index from the array
// determine if that index contains a "paragraph" i.e. by length or something
// return the text from that paragraph and its index in the array
// index is important for finding chapter later
// probably need a paragraph struct to hold text and location

func GetParagraph(text []string) (paragraph []string) {
  var graph []string
  for i, line := range text {
    if i < 20 {
      graph = append(graph, line)
    }
    paragraph = graph[:]
  }
  return
}

// outputs a line-by-line copy of the text in the book file
func LinebyLineScan(scanner *bufio.Scanner) {
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
