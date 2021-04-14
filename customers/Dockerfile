# Builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o store-customers .


# Generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/store-customers .

EXPOSE 8080

# Executable
ENTRYPOINT [ "./store-customers" ]
