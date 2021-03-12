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
  fmt.Printf("%v\n", GetParagraph(book.bookText))
  PrintBook(book.bookText)
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
  bookText []string
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
  fullText := Strip(GetTextAll(scanner))
  bookText := GetText(fullText)
  return Book{filename, title, author, date, fullText, bookText}

}

func BookPrinter(book Book) {
  fmt.Println("Filename: ", book.filename)
  fmt.Println("Title: \t\t", book.title)
  fmt.Println("Author: \t", book.author)
  fmt.Println("Release Date: \t", book.date)
  //fmt.Println("Full Text: \n", book.fullText)
  fmt.Println("Book Text: \n", book.bookText)
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

// return all text because I couldn't run range over a scanner type
func GetTextAll(scanner *bufio.Scanner) (text []string) {
  for scanner.Scan() {
    text = append(text, scanner.Text())
  }
  return
}

// finds the new lines and returns a splice of paragraphs
func GetText(fulltext []string) (text []string) {
  for _, line := range fulltext {
    switch line {
    case "":
      line = "NEWLINE"
      text = append(text, line)
    default: 
      text = append(text, line)
    }
  }
  textstring := strings.TrimSpace(strings.Join(text[:], " "))
  text = strings.Split(textstring, "NEWLINE")
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
    if i < 10 {
      graph = append(graph, line)
    }
    paragraph = graph[:]
  }
  fmt.Println("Book Length: ", len(text))
  fmt.Println(len(paragraph))
  return
}

// paragraph by paragraph printing of book
func PrintBook(text []string) {
  for i, paragraph := range text {
    fmt.Println(i, paragraph)
  }
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
