FROM postgres:9.3

ENV POSTGRES_USER docker
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB docker
ADD ./db/postgres/schema/schema.sql /docker-entrypoint-initdb.d/
