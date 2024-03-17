# Stage 1 - Build Vue application
FROM node:lts-alpine as build-stage

WORKDIR /app

COPY ./app/package*.json ./

RUN npm install

COPY ./app/ .

RUN npm run build

# Stage 2 - Build Go backend
FROM golang:1.22

WORKDIR /usr/src/app

COPY --from=build-stage /app/dist ./dist

COPY ./server/go.mod ./server/go.sum ./

RUN go mod download && go mod verify

COPY ./server/ .

RUN go build -v -o /usr/local/bin/app/ ./...

EXPOSE 8080

CMD ["/usr/local/bin/app/notes"]