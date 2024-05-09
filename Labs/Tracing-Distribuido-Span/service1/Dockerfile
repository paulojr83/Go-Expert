FROM golang:latest as build

WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o api cmd/main.go

FROM scratch
ENV SERVICE1_API="http://localhost:8181/cep/%s"
ENV OTEL_SERVICE_NAME="Service1"
ENV HTTP_PORT=":8080"

COPY --from=build /app/api ./
CMD ["./api"]