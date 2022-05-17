FROM golang:latest as BUILDER

ARG GOPROXY
ARG GONOSUMDB

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o server ./main.go
FROM golang:latest

COPY --from=BUILDER /app/server /app/server

EXPOSE 8080
ENTRYPOINT /app/server
