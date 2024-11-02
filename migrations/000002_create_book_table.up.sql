CREATE TABLE IF NOT EXISTS publisher (
    publisher_id SERIAL PRIMARY KEY,
    publisher_name VARCHAR(127) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS category
(
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(63) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS author
(
    author_id SERIAL PRIMARY KEY,
    author_name VARCHAR(63) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS book
(
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(127) NOT NULL,
    category_id INT NOT NULL REFERENCES category (category_id),
    description TEXT,
    publisher_id INT NOT NULL REFERENCES publisher (publisher_id),
    year_publishing INT NOT NULL,
    isbn VARCHAR(63),
    amount INT DEFAULT 0
    -- FOREIGN KEY (category_id) REFERENCES category (category_id),
    -- FOREIGN KEY (publisher_id) REFERENCES publisher (publisher_id)
);

CREATE TABLE IF NOT EXISTS author_book
(
    author_book_id SERIAL PRIMARY KEY,
    author_id INT NOT NULL REFERENCES author,
    book_id INT NOT NULL REFERENCES book,
    UNIQUE (author_id, book_id)
);
