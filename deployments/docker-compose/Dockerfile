FROM golang:1.14.1

# if you want to change localtime, uncomment below.
# RUN rm /etc/localtime && echo Asia/Tokyo > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata

WORKDIR /go/src/github.com/inflion/inflion
COPY . .
ENV GO111MODULE=on
RUN go get -u github.com/cosmtrek/air

ENTRYPOINT ["air"]
