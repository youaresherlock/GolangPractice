/*
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}
*/
package main

import "fmt"

type Books struct {
	title string 
	author string
	subject string
	book_id int
}

func main() {
	/*Declare Book1 of type Book*/
	var Book1 Books
	var Book2 Books

	/*book 1 specification*/
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	printBook(&Book1)

	printBook(&Book2)
}

/*要使用指向该结构体的指针来访问结构的成员,必须使用.号*/
func printBook(book *Books) {
	fmt.Println("Book title: ", book.title)
	fmt.Println("Book author: ", book.author)
	fmt.Println("Book subject: ", book.subject)
	fmt.Println("Book book_id: ", book.book_id)
}