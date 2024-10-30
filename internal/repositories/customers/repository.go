package customers

import (
	"db-coursework/internal/models"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	// size of customers batch to insert per query
	batch = 100
)

// * idk should i use squirrel or not
type repository struct {
	db *sqlx.DB
	// builder sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
		// builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *repository) AddCustomers(customers []models.Customer) ([]int, error) {
	createdCustomersIDs := make([]int, 0, len(customers))
	for i := 0; i < len(customers); i += batch {
		endOfBatch := i + batch
		if endOfBatch > len(customers) {
			endOfBatch = len(customers)
		}

		newIDs, err := r.addCustomers(customers[i:endOfBatch])
		if err != nil {
			return nil, errors.Wrap(err, "error during adding customers batch")
		}

		createdCustomersIDs = append(createdCustomersIDs, newIDs...)
	}

	return createdCustomersIDs, nil
}

// TODO: maybe i can refactor it to return created []models.Customer
// TODO: but before this return statement i should add id to passed []models.Customer as argument
func (r *repository) addCustomers(customers []models.Customer) ([]int, error) {
	if len(customers) == 0 {
		return nil, nil
	}

	valuesQuery := make([]string, 0, len(customers))
	valuesArgs := make([]interface{}, 0, len(customers))
	i := 1
	for _, customer := range customers {
		valuesQuery = append(valuesQuery,
			fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i, i+1, i+2, i+3, i+4, i+5, i+6, i+7),
		)
		i += 8
		valuesArgs = append(
			valuesArgs, customer.LastName, customer.FirstName, customer.FatherName, customer.Gender,
			customer.DateofBirth, customer.Phone, customer.Email, customer.Address,
		)
	}

	query := fmt.Sprintf(`
		INSERT INTO customer (last_name, first_name, father_name, gender, date_of_birth, phone, email, address)
		VALUES %s
		RETURNING customer_id
	`, strings.Join(valuesQuery, ", "))

	rows, err := r.db.Query(query, valuesArgs...)
	if err != nil {
		return nil, errors.Wrap(err, "error during customer insertion")
	}

	newIDs := make([]int, 0, len(customers))
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "error getting new customer id")
		}
		newIDs = append(newIDs, id)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error getting ids of new custoemrs")
	}

	return newIDs, nil
}
