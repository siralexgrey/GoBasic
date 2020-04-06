package main

import "fmt"

type EDocument struct {
	Number        string
	Date          string
	NumberOfPages int
	// Footer -- анонимное поле сразу определенное типом структуры
	FooterDoc Footer // именованное поле с типом
}

type Footer struct {
	Author string
}

func (doc EDocument) showAuthor() {
	fmt.Println(doc.FooterDoc.Author)
}

// Использование указателя передаёт в метод сам объект, а него копию
// Как следсвия - изменения просиходят в самом объекте
// Так же сокращает потребление памяти за счет того, что не создается копия объекта
func (doc *EDocument) showNumberOfPages() {
	fmt.Println(doc.NumberOfPages)
}

func (doc *EDocument) clearAllPages() {
	doc.NumberOfPages = 0
	fmt.Println(doc.NumberOfPages)
}

func (doc *EDocument) setPrefixAuthore() {
	doc.FooterDoc.Author = "the Author is " + doc.FooterDoc.Author
	fmt.Println(doc.FooterDoc.Author)
}

func main() {

	doc1 := EDocument{
		Number:        "001-A",
		Date:          "06.04.2020",
		NumberOfPages: 10,
		FooterDoc: Footer{
			Author: "Author-1",
		},
	}

	var doc2 EDocument
	doc2.Number = "002-A"
	doc2.Date = "06.04.2020"
	doc2.NumberOfPages = 5
	doc2.FooterDoc.Author = "Author-2"

	doc3 := new(EDocument)
	doc3.Number = "003-B"
	doc3.Date = "06.04.2020"
	doc3.NumberOfPages = 7
	doc3.FooterDoc.Author = "No name"

	// fmt.Printf("%T - %v\n", doc1, doc1)
	// fmt.Printf("%T - %v\n", doc2, doc2)
	// fmt.Printf("%T - %v\n", doc3, doc3)
	// fmt.Printf("%T - %v\n", doc2.FooterDoc, doc2.FooterDoc)

	// doc1.showAuthor()
	// doc2.showAuthor()
	// doc3.showAuthor()
	doc1.clearAllPages()
	fmt.Println(doc1)
}
