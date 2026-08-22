CREATE SCHEMA konvert;

CREATE TABLE konvert.users (
  id SERIAL PRIMARY KEY,
  varsion BIGINT NOT NULL DEFAULT 1,
  full_name VARCHAR(100) NOT NULL CHECK(char_length(full_name) BETWEEN 2 AND 100),
  phone_number VARCHAR(15) CHECK(
    phone_number ~ '^\+[0-9]+$'
    AND
    char_length(phone_number) BETWEEN 10 AND 15
  )
);


