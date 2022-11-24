postgres:
	docker run --name postgres12 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_trx

dropdb:
	docker exec -it postgres12 dropdb simple_trx

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5431/simple_trx?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5431/simple_trx?sslmode=disable" -verbose down

sqlc-generate:
	docker run --rm -v "C:\Users\Mahatma Ageng Wisesa\Desktop\Unit Test\go\DB-Migration\be-4-rd:/src" -w /src kjconroy/sqlc generate

sqlc-init:
	docker run --rm -v "C:\Users\Mahatma Ageng Wisesa\Desktop\Unit Test\go\DB-Migration\be-4-rd:/src" -w /src kjconroy/sqlc init

go-test-cover:
	go test -covermode=count -coverpkg=./... -coverprofile cover.out -v ./...
	go tool cover -html cover.out -o cover.html

server:
	go run main.go

mockgen:
	mockgen -package mockdb -destination src/db/mock/store.go be-4-rs-crud/src/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc-generate sqlc-init go-test-cover server mockgen