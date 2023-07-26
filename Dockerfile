FROM golang:alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:alpine AS build
WORKDIR /build
COPY --from=modules /go/pkg /go/pkg
COPY . .
RUN go build -o shrtnr cmd/app/main.go

FROM alpine
WORKDIR /app
COPY --from=build /build/shrtnr .
CMD ["./shrtnr"]