-- public.account definition
-- Drop table
DROP TABLE IF EXISTS account;
CREATE TABLE account (
    id serial PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    city VARCHAR,
    country VARCHAR,
    state VARCHAR,
    avaliable_features VARCHAR,
    cannabis_repository_id INTEGER,
    description VARCHAR(255),
    language VARCHAR,
    status VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);