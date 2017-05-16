package main
import "fmt"
type Books struct {
  title string
  author string
  subject string
  book_id int
}
func main()  {
  var Book1 Books
  var Book2 Books
  Book1.title = "go 语言"
  Book1.author = "go 语言作者"
  Book1.subject = "go 语言教程"
  Book1.book_id = 123456

  Book2.title = "Pyhton 语言"
  Book2.author = "Pyhton 语言作者"
  Book2.subject = "Pyhton 语言教程"
  Book2.book_id = 12312312
  printBook(&Book1)
  printBook(&Book2)
}
func printBook(book *Books)  {
  fmt.Printf("book title %s\n",book.title);
  fmt.Printf("book author %s\n",book.author);
  fmt.Printf("book subject %s\n",book.subject);
  fmt.Printf("book id%d\n",book.book_id);
}
