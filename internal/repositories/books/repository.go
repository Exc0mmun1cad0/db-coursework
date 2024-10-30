// Not only books lies there but also connected entities such as authors, categories, publishers
package books

import (
	"db-coursework/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
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

func (r *repository) AddPublisher(publisher models.Publisher) (uint64, error) {
	var publisherID uint64
	err := r.db.Get(
		&publisherID, `
			INSERT INTO publisher(publisher_name)
			VALUES ($1)
			RETURNING publisher_id
			ON CONFLICT DO NOTHING;
		`, publisher.Name,
	)
	if err != nil {
		return 0, fmt.Errorf("error during publisher insert")
	}

	return publisherID, nil
}

func (r *repository) AddCategory(category models.Category) (uint64, error) {
	var categoryID uint64
	err := r.db.Get(
		&categoryID, `
			INSERT INTO category(category_name)
			VALUES ($1)
			RETURNING category_id
			ON CONFLICT DO NOTHING;
		`, category.Name,
	)
	if err != nil {
		return 0, fmt.Errorf("error during category insert")
	}

	return categoryID, nil
}

func (r *repository) AddAuthor(author models.Author) (uint64, error) {
	var authorID uint64
	err := r.db.Get(
		&authorID, `
			INSERT INTO category(category_name)
			VALUES ($1)
			RETURNING category_id
			ON CONFLICT DO NOTHING;
		`, author.Name,
	)
	if err != nil {
		return 0, fmt.Errorf("error during author insert")
	}

	return authorID, nil
}
