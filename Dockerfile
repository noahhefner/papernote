# Build Go Application

FROM golang:1.22

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app/ ./...

# Make default db and notes directories
RUN mkdir -p /data/db
RUN mkdir -p /data/notes

EXPOSE 8080

CMD ["/usr/local/bin/app/notes"]
