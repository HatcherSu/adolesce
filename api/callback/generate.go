package api

//go:generate protoc --proto_path=../../third_party --proto_path=. --go_out=. --go_opt=paths=source_relative --go-http_out=paths=source_relative:. *.proto
