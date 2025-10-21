FROM golang:latest as build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o task-manager .

FROM debian:stable-slim

WORKDIR /app
COPY --from=build-stage /app/task-manager .

COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

EXPOSE 8080

CMD ["./wait-for-it.sh", "db:3306", "--timeout=60", "--", "./task-manager"]