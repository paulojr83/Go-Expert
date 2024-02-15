module github.com/paulojr83/Go-Expert/packaging/sistema

go 1.21.5

replace github.com/paulojr83/Go-Expert/packaging/math => ../math

//go mod edit -replace github.com/paulojr83/Go-Expert/packaging/math=../math

require github.com/paulojr83/Go-Expert/packaging/math v0.0.0-00010101000000-000000000000
