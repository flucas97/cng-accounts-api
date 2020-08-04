# docker build . --build-arg app_env=dev -t cng-account
# docker run -it cng-account 
FROM golang:latest
LABEL maintainer="flucas97 <feijolucas1997gmail.com>"

ARG app_env
ENV APP_ENV $app_env

WORKDIR /app/src/accounts

ENV GOPATH=/app
COPY . /app/src/accounts

RUN go mod download

RUN go get ./
RUN go build main.go

ENTRYPOINT [ "./main" ]
EXPOSE 8081
