CREATE TYPE gender AS ENUM('Мужчина', 'Женщина');
CREATE TABLE IF NOT EXISTS customer
(
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(64) NOT NULL,
    first_name VARCHAR(64) NOT NULL,
    father_name VARCHAR(64),
    gender gender,
    date_of_birth DATE NOT NULL,
    phone VARCHAR(32),
    email VARCHAR(64),
    address VARCHAR(128) NOT NULL
);
