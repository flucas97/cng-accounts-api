FROM golang:latest
LABEL maintainer="flucas97 <feijolucas19972gmail.com>"

WORKDIR /account
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

EXPOSE 8081
CMD ["./main"]