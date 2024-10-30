package models

type Author struct {
	ID   uint64
	Name string
}

type Category struct {
	ID   uint64
	Name string
}

type Publisher struct {
	ID   uint64
	Name string
}

type Book struct {
	ID            uint64
	Title         string
	Authors       []Author
	Category      Category
	Description   string
	YearPublished int64
	Publishers    []Publisher
	ISBN          string
	Amount        uint64
}
