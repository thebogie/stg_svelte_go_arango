# builder image
FROM golang:alpine3.15 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main server.go


# generate clean, final image for end users
FROM alpine:3.15
COPY --from=builder /build/main .

# executable
ENTRYPOINT [ "./main" ]
