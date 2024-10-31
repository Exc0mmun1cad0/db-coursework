// Not only books lies there but also connected entities such as authors, categories, publishers
package books

import (
	"db-coursework/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	// size of customers batch to insert per query
	batch = 100
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) addAttribute(attribute models.BookAttribute, name string) (uint64, error) {
	var attributeID uint64

	query := fmt.Sprintf(`
		INSERT INTO %s(%s_name) VALUES($1)
		RETURNING %s_id
		ON CONFLICT DO NOTHING;
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
