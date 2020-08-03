# docker build . --build-arg app_env=dev -t cng-account
# docker run -it cng-account 
FROM golang:latest
LABEL maintainer="flucas97 <feijolucas1997gmail.com>"

ARG app_env
ENV APP_ENV $app_env

WORKDIR /account
COPY go.mod ./
RUN go mod download
COPY . .
RUN go get ./
RUN go build 

CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 8081
