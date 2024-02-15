module github.com/paulojr83/Go-Expert/packaging/sistema

go 1.21.5

//go work init ./math ./sistema

require github.com/google/uuid v1.5.0

//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out
//go test -bench=.
//go test -bench=. -run=^#
//go test -bench=. -run=^# -count=10
//go test -bench=. -run=^# -count=10 -benchtime=3s -benchmen
