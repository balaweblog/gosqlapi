FROM golang:1.6

RUN go get 	github.com/go-sql-driver/mysql

EXPOSE 8080
