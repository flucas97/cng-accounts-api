FROM postgres:11

COPY ./create_db.sh /docker-entrypoint-initdb.d/20-create_db.sh
COPY /db/postgres/schema/schema.sql /db/postgres/schema/schema.sql

RUN chmod +x /docker-entrypoint-initdb.d/20-create_db.sh