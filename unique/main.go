package main

import (
	"fmt"
	"unique"
)

func main() {
	{
		empty1 := unique.Make("")
		empty2 := unique.Make("")
		value := unique.Make("abcdefg")

		fmt.Printf("empty==empty: %v\n", empty1 == empty2)
		fmt.Printf("empty==value: %v\n", empty1 == value)
	}
	{
		var (
			zeroNonMake unique.Handle[User]
			zero        = unique.Make(User{})
		)
		fmt.Printf("zerot==zero: %v\n", zeroNonMake == zero)

		u1l := unique.Make(User{ID: 1, Name: "Alice"})
		u1r := unique.Make(User{ID: 1, Name: "Alice"})
		fmt.Printf("u1l==u1r: %v\n", u1l == u1r) // -> true

		u1lp := unique.Make(&User{ID: 1, Name: "Alice"})
		u1rp := unique.Make(&User{ID: 1, Name: "Alice"})
		fmt.Printf("u1lp==u1rp: %v\n", u1lp == u1rp) // -> false
	}
	{
		book1 := NewBook("magic&magic", "alice", "Fantasy", "Magic")
		showBookType(book1)

		book2 := NewBook("idol&idol", "bob", "Idol", "Idol")
		showBookType(book2)

		book3 := NewBook("unknown&unknown", "charlie", "Unknown", "Unknown")
		showBookType(book3)
	}
}

type User struct {
	ID   int
	Name string
}

func NewBook(title, autor, category, subCategory string) Book {
	typ := BookType{
		Category:    category,
		SubCategory: subCategory,
	}
	return Book{
		Title:      title,
		Autor:      autor,
		Type:       typ,
		TypeUnique: unique.Make(typ),
	}
}

var (
	magicTyp = unique.Make(BookType{"Fantasy", "Magic"})
	idolTyp  = unique.Make(BookType{"Idol", "Idol"})
)

func showBookType(book Book) {
	switch book.TypeUnique {
	case magicTyp:
		fmt.Println("This is a magic book")
	case idolTyp:
		fmt.Println("This is an idol book")
	default:
		fmt.Println("This is an unknown book")
	}
}

type BookType struct {
	Category    string
	SubCategory string
}

type Book struct {
	ID         int
	Title      string
	Autor      string
	Type       BookType
	TypeUnique unique.Handle[BookType]
}
