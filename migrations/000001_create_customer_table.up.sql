CREATE TYPE gender AS ENUM('male', 'female');
CREATE TABLE IF NOT EXISTS customer
(
    customer_id SERIAL PRIMARY KEY,
    last_name VARCHAR(64) NOT NULL,
    first_name VARCHAR(64) NOT NULL,
    father_name VARCHAR(64),
    gender gender,
    birth_date DATE NOT NULL,
    phone VARCHAR(32),
    email VARCHAR(64),
    address VARCHAR(128) NOT NULL
);
