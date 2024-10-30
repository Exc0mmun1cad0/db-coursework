CREATE TABLE IF NOT EXISTS publisher (
    publisher_id SERIAL PRIMARY KEY,
    publisher_name VARCHAR(127) NOT NULL
);

CREATE TABLE IF NOT EXISTS category
(
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(63) NOT NULL
);

CREATE TABLE IF NOT EXISTS author
(
    author_id SERIAL PRIMARY KEY,
    author_name VARCHAR(63) NOT NULL
);

CREATE TABLE IF NOT EXISTS book
(
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(127) NOT NULL,
    category_id INT NOT NULL,
    description TEXT,
    year_publishing INT DEFAULT 0,
    isbn VARCHAR(31),
    amount INT DEFAULT 0,
    FOREIGN KEY (category_id) REFERENCES category (category_id)
);

CREATE TABLE IF NOT EXISTS author_book
(
    author_book_id SERIAL PRIMARY KEY,
    author_id INT NOT NULL REFERENCES author,
    book_id INT NOT NULL REFERENCES book,
    UNIQUE (author_id, book_id)
);

CREATE TABLE IF NOT EXISTS publisher_book
(
    publisher_book_id SERIAL PRIMARY KEY,
    publisher_id INT NOT NULL REFERENCES publisher (publisher_id),
    book_id INT NOT NULL REFERENCES book (book_id),
    UNIQUE (publisher_id, book_id)
);
