FROM golang:latest as build

WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server cmd/main.go

FROM scratch
COPY --from=build /app/server ./
CMD ["./server"]