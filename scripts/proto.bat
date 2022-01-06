cd C:\Users\apava\go\src\github.com\apavanello\CALCULATOR
protoc .\proto\*.proto --go_out=plugins=grpc:.
go mod tidy