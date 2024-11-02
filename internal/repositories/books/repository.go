// Not only books lies there but also connected entities such as authors, categories, publishers
package books

import (
	"db-coursework/internal/models"
	"fmt"
	"log"
	"math/rand"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

// This function is runned only once when filling db with books data
func (r *repository) AddBooks(books []models.Book) ([]uint64, error) {
	bookIDs := make([]uint64, 0, len(books))

	authorMap := make(map[string]uint64)
	publisherMap := make(map[string]uint64)
	categoryMap := make(map[string]uint64)

	for i, book := range books {
		log.Printf("inserting book: %d", i+1)

		var err error
		// Publisher
		publisherID, ok := publisherMap[book.Publisher.Name]
		if !ok {
			publisherID, err = r.AddPublisher(book.Publisher)
			publisherMap[book.Publisher.Name] = publisherID
			if err != nil {
				return nil, errors.Wrap(err, "error during insertion publisher into repository")
			}
		}
		// Category
		categoryID, ok := categoryMap[book.Category.Name]
		if !ok {
			categoryID, err = r.AddCategory(book.Category)
			categoryMap[book.Category.Name] = categoryID
			if err != nil {
				return nil, errors.Wrap(err, "error during insertion category into repository")
			}
		}
		bookID, err := r.addBook(book.Title, book.Description, book.ISBN, categoryID, publisherID, book.YearPublishing)
		if err != nil {
			return nil, errors.Wrap(err, "error during book insert")
		}
		bookIDs = append(bookIDs, bookID)

		// Authors
		for _, author := range book.Authors {
			authorID, ok := authorMap[author.Name]
			if !ok && authorID == 0 {
				authorID, err = r.AddAuthor(author)
				authorMap[author.Name] = authorID
				if err != nil {
					return nil, errors.Wrap(err, "error during insertion author into repository")
				}
			}
			_, err := r.addAuthorBook(bookID, authorID)
			if err != nil {
				return nil, errors.Wrap(err, "error during insertion new author book relation")
			}
		}

	}

	return bookIDs, nil
}

func (r *repository) addBook(title, description, isbn string, category_id, publisher_id uint64, year_publishing int64) (uint64, error) {
	var bookID uint64
	amount := rand.Intn(15) + 1

	query := `INSERT INTO book(title, category_id, description, publisher_id, year_publishing, isbn, amount)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING book_id
	`

	err := r.db.Get(&bookID, query, title, category_id, description, publisher_id, year_publishing, isbn, amount)
	if err != nil {
		return 0, errors.Wrap(err, "error during book insert")
	}

	return bookID, nil
}

func (r *repository) addAuthorBook(bookID, authorID uint64) (uint64, error) {
	var bookAuthorID uint64

	query := `INSERT INTO author_book(author_id, book_id) VALUES ($1, $2) RETURNING author_book_id`

	err := r.db.Get(&bookAuthorID, query, authorID, bookID)
	if err != nil {
		return 0, errors.Wrap(err, "error during adding new author book relation")
	}

	return bookAuthorID, nil
}

type bookAttribute interface {
	GetName() string
}

func (r *repository) addAttribute(attribute bookAttribute, name string) (uint64, error) {
	var attributeID uint64

	query := fmt.Sprintf(`
		INSERT INTO %s(%s_name) VALUES($1)
		RETURNING %s_id;
	`, name, name, name)

	err := r.db.Get(&attributeID, query, attribute.GetName())
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("error during %s insert", name))
	}

	return attributeID, nil
}

func (r *repository) AddPublisher(publisher models.Publisher) (uint64, error) {
	return r.addAttribute(publisher, "publisher")
}

func (r *repository) AddCategory(category models.Category) (uint64, error) {
	return r.addAttribute(category, "category")
}

func (r *repository) AddAuthor(author models.Author) (uint64, error) {
	return r.addAttribute(author, "author")
}
