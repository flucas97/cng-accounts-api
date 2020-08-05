-- public.account definition
-- Drop table
DROP TABLE IF EXISTS account;
CREATE TABLE account (
    id serial primary key,
    name text not null unique,
    email text not null unique,
    password text not null,
    city text,
    country text,
    state text,
    avaliable_features text,
    cannabis_repository_id int,
    description varchar(255),
    language text,
    status text not null,
    updated_at timestamp NOT NULL DEFAULT now(),
    created_at timestamp NOT NULL DEFAULT now()
);