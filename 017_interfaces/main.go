package main

import "fmt"

/*
Полиморфизм

1 - Возможность объектов с одинаковой спецификацией иметь различную реализацию

2 - Способность объекта использовать методы производного класса
	который не существует на момент создания базового

3 - Это свойство классов иметь одинаковые методы,
	которые будут работать по-разному в контексте объектов
*/

type Document struct {
	Date          string
	Number        string
	NumberOfPages int
}

type PersonCard struct {
	Date      string
	FirstName string
	LastName  string
	Age       int
}

type PrintInterface interface {
	CheckData()
}

// func (d *Documnet) PritD() {...}

func (d *Document) CheckData() {
	if d.Date != "" {
		fmt.Println("Date in the doc is correct")
	} else {
		fmt.Println("Date in the doc is incorrect")
	}
}

func (p *PersonCard) CheckData() {
	if p.Date != "" {
		fmt.Println("Date in the card is correct")
	} else {
		fmt.Println("Date in the card is incorrect")
	}
}

func main() {

	doc := new(Document)
	pcard := new(PersonCard)

	doc.Date = "06.04.2020"
	doc.NumberOfPages = 5
	doc.Number = "A - 100"

	pcard.Date = "06.04.2020"
	pcard.Age = 21
	pcard.FirstName = "User"
	pcard.LastName = "Test"

	sl := []PrintInterface{doc, pcard}

	PrintAnyDoc(sl)
}

func PrintAnyDoc(anyDoc []PrintInterface) {
	for _, v := range anyDoc {
		fmt.Println(v)
	}
}
