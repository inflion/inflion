# build stage
FROM golang:1.14.2-alpine AS build-env
COPY . /src
RUN cd /src && go build -o run-server

# final stage
FROM alpine
RUN echo $TZ > /etc/timezone
WORKDIR /app
COPY --from=build-env /src/run-server /app/
RUN ls /app/
ENTRYPOINT /app/run-server