CREATE TABLE IF NOT EXISTS book_loan
(
    book_loan_id SERIAL PRIMARY KEY, 
    book_id INT NOT NULL REFERENCES book (book_id),
    customer_id INT NOT NULL REFERENCES customer (customer_id),
    date_loaned DATE NOT NULL,
    date_due DATE NOT NULL,
    date_returned DATE,
    amount INT CHECK (amount > 0)
);
