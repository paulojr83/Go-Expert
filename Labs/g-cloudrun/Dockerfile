FROM golang:latest as build

WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o api cmd/main.go

FROM scratch
COPY --from=build /app/api ./
CMD ["./api"]