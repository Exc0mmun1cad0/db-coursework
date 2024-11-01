package models

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

func (a Author) GetName() string {
	return a.Name
}

func (c Category) GetName() string {
	return c.Name
}

func (p Publisher) GetName() string {
	return p.Name
}
