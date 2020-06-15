FROM golang:1.14.2-alpine3.11
RUN apk --update-cache --no-cache add git build-base tzdata
RUN go get -v github.com/rubenv/sql-migrate/sql-migrate
WORKDIR /go/src/github.com/inflion/inflion
ENTRYPOINT ["/go/bin/sql-migrate"]