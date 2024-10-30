CREATE TYPE gender AS ENUM('Мужчина', 'Женщина');
CREATE TABLE IF NOT EXISTS customer
(
    customer_id SERIAL PRIMARY KEY,
    last_name VARCHAR(63) NOT NULL,
    first_name VARCHAR(63) NOT NULL,
    father_name VARCHAR(63),
    gender gender,
    date_of_birth DATE NOT NULL,
    phone VARCHAR(31),
    email VARCHAR(63),
    address VARCHAR(255) NOT NULL
);
