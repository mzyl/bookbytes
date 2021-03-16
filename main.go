package main

import (
  "os"
  "fmt"
  "log"
  "time"
  "bufio"
  "strings"
  "math/rand"
)

func main() {
  //batching()
  filename := "./books/36-0.txt"
  book := BookBuilder(filename)
  BookPrinter(book)
  fmt.Println()
  book.paragraph = GetParagraph(book)
  fmt.Println("Index: ", book.paragraph)
  fmt.Println(ParagraphPrinter(book))
  fmt.Println()
  book.paragraph = NextParagraph(book)
  fmt.Println("Index: ", book.paragraph)
  fmt.Println(ParagraphPrinter(book))
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
  paragraph int
}

/*** BOOKBUILDER FUNCTIONS ***/


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
  fullText := StripLicense(GetAllText(scanner))
  bookText := GetBookText(fullText)
  paragraph := 0
  return Book{filename, title, author, date, fullText, bookText, paragraph}
}

// prints all contents of Book structure
func BookPrinter(book Book) {
  fmt.Println("Filename: ", book.filename)
  fmt.Println("Title: \t\t", book.title)
  fmt.Println("Author: \t", book.author)
  fmt.Println("Release Date: \t", book.date)
  //fmt.Println("Full Text: \n", book.fullText)
  //fmt.Println("Book Text: \n", book.bookText)
}

func ParagraphPrinter(book Book) string {
  return book.bookText[book.paragraph]
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

// return all text because I couldn't run range over a scanner type
func GetAllText(scanner *bufio.Scanner) (text []string) {
  for scanner.Scan() {
    text = append(text, scanner.Text())
  }
  return
}

// finds the new lines and returns a splice of paragraphs
func GetBookText(fulltext []string) (text []string) {
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
func StripLicense(text []string) (stripped []string) {
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
// navigating up and down can result in blank paragraphs since not all empty strings are being removed

/*** PARAGRAPH/NAVIGATION FUNCTIONS ***/


func GetParagraph(book Book) (index int) {
  text := book.bookText
  rand.Seed(time.Now().UnixNano())
  var randomparagraph int
  for range text {
    randomparagraph = rand.Intn(len(text))
    if len(text[randomparagraph]) > 400 {
      index = randomparagraph
      break
    }
  }
  return
}

// is there ever a reason be go forward, then to the paragraph above the original?

func PreviousParagraph(book Book) int {
  return book.paragraph - 1
}

func NextParagraph(book Book) int {
  return book.paragraph + 1
}

func BeginningChapter(book Book) int {
  // find index of paragraph at the beginning of the current chapter
  return 0
}

func BeginningBook(book Book) int {
  // find index of paragraph at the beginning of the book
  // should this just return full book?
  return 0
}

// any reason to return the entire book for the user to scroll through?


/*** HELPER FUNCTIONS ***/


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
