# Build Go Application

FROM golang:1.22

WORKDIR /usr/src/app

COPY ./server/go.mod ./server/go.sum ./

RUN go mod download && go mod verify

COPY ./server/ .

RUN go build -v -o /usr/local/bin/app/ ./...

EXPOSE 8080

CMD ["/usr/local/bin/app/notes"]