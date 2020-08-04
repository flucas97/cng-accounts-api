CREATE TABLE accounts (
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
ALTER TABLE accounts OWNER TO master;
GRANT ALL ON TABLE accounts TO master;