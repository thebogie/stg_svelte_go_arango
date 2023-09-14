# builder image
FROM golang:1.21

WORKDIR /stg

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN  CGO_ENABLED=0 GOOS=linux go build -a -o runbackend.go  ./app/server.go

EXPOSE 50002

CMD ["./runbackend.go"]
