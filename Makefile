postgres:
	docker compose up

createdb: 
	docker exec -it some-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it some-postgres dropdb simple_bank
 
migrateup: 
	migrate -path db/migration -database "postgresql://root:secretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

dbversion:
	migrate -path db/migration -database "postgresql://root:secretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose version

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc