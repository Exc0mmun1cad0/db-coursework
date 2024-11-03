package models

type Book struct {
	ID             uint64
	Title          string
	Authors        []Author
	Category       Category
	Description    string
	YearPublishing int64
	Publisher      Publisher
	Amount         uint64
	ISBN           string
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
