FROM golang:alpine AS build
WORKDIR /build
COPY . .
RUN go build -o shrtnr cmd/app/main.go

FROM alpine
WORKDIR /app
COPY --from=build /build/shrtnr .
CMD ["./shrtnr"]
EXPOSE 8080