-- public.account definition
-- Drop table
-- DROP TABLE public.account;
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

CREATE OR REPLACE VIEW metrics
  AS SELECT
    account.id,
    account.name,
    account.email,
    account.password,
    account.city,
    account.country,
    account.state,
    account.avaliable_features,
    account.cannabis_repository_id,
    account.description,
    account.language,
    account.status,
    account.updated_at,
    account.created_at
   FROM account;

ALTER TABLE account OWNER TO postgres;
GRANT ALL ON TABLE account TO postgres;
