CREATE SCHEMA konvert;

CREATE TABLE konvert.users (
  id SERIAL PRIMARY KEY,
  version BIGINT NOT NULL DEFAULT 1,
  full_name VARCHAR(100) NOT NULL CHECK(char_length(full_name) BETWEEN 2 AND 100),
  phone_number VARCHAR(15) CHECK(
    phone_number ~ '^\+[0-9]+$'
    AND
    char_length(phone_number) BETWEEN 10 AND 15
  ),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE konvert.receipts (
  id SERIAL PRIMARY KEY,
  version BIGINT NOT NULL DEFAULT 1,
  user_id INT NOT NULL REFERENCES konvert.users(id) ON DELETE CASCADE,
  name VARCHAR(100) NOT NULL CHECK(char_length(name) BETWEEN 2 AND 100),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  amount NUMERIC(10, 2) NOT NULL CHECK(amount >= 0)
);

CREATE TABLE konvert.covers (
  id SERIAL PRIMARY KEY,
  version BIGINT NOT NULL DEFAULT 1,
  name VARCHAR(255) NOT NULL CHECK(char_length(name) BETWEEN 2 AND 255),
  user_id INT NOT NULL REFERENCES konvert.users(id) ON DELETE CASCADE,
  max_month_amount NUMERIC(10, 2) NOT NULL CHECK(max_month_amount >= 0),
  color varchar(7),
  icon varchar(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE konvert.expenses (
  id SERIAL PRIMARY KEY,
  version BIGINT NOT NULL DEFAULT 1,
  name VARCHAR(255) NOT NULL CHECK(char_length(name) BETWEEN 2 AND 255),
  amount NUMERIC(10, 2) NOT NULL CHECK(amount >= 0),
  receipt_id INT NOT NULL REFERENCES konvert.receipts(id) ON DELETE CASCADE,
  user_id INT NOT NULL REFERENCES konvert.users(id) ON DELETE CASCADE,
  cover_id INT NOT NULL REFERENCES konvert.covers(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);