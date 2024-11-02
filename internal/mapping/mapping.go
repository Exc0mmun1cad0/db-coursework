// This package contains functions for mapping obejcts which are received from chitaigorod API to models such as books, authors and publishers
package mapping

import (
	chitaigorod "db-coursework/internal/api/chitai_gorod"
	"db-coursework/internal/models"
	"strings"

	"github.com/pkg/errors"
)

func ResponseToModel(books []chitaigorod.Data) ([]models.Book, error) {
	modelBooks := make([]models.Book, 0, len(books))
	for _, book := range books {
		modelBook, err := responseToModel(&book)
		if err != nil {
			return nil, errors.Wrap(err, "cannot convert chitaigorod api response to models")
		}
		modelBooks = append(modelBooks, *modelBook)
	}

	return modelBooks, nil
}

func responseToModel(book *chitaigorod.Data) (*models.Book, error) {
	bookModel := &models.Book{}

	bookModel.Title = book.Attributes.Title
	bookModel.Category = models.Category{
		ID:   uint64(book.Attributes.ID),
		Name: book.Attributes.Category.Title,
	}
	bookModel.Description = book.Attributes.Description
	bookModel.YearPublishing = int64(book.Attributes.YearPublishing)
	bookModel.Publisher = models.Publisher{
		ID:   uint64(book.Attributes.ID),
		Name: book.Attributes.Publisher.Title,
	}
	bookModel.Authors = make([]models.Author, 0, len(book.Attributes.Authors))
	for _, author := range book.Attributes.Authors {
		bookModel.Authors = append(bookModel.Authors,
			models.Author{
				ID:   uint64(author.ID),
				Name: strings.Join([]string{author.LastName, author.MiddleName, author.FirstName}, " "),
			},
		)
	}

	isbn, err := chitaigorod.GetISBN(book.Attributes.URL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get book isbn")
	}
	bookModel.ISBN = isbn

	return bookModel, nil
}
