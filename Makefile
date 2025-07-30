a:
	docker ps -a
stop:
	docker stop postgres14.18
start:
	docker start postgres14.18
postgres:
	docker run --name postgres14.18 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.18
createdb:
	docker exec -it postgres14.18 createdb --username=root --owner=root simplebank
dropdb:
	docker exec -it postgres14.18 dropdb simplebank
db:
	docker exec -it postgres14.18 psql -U root -d simplebank
up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up
down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v ./db/sqlc/... -cover

.PHONY: postgres createdb dropdb up down sqlc test