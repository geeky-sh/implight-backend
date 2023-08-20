migrate_up:
	migrate -path db/migrations -database "postgres://aash:@localhost:5432/implight?sslmode=disable" -verbose up
migrate_down:
	migrate -path db/migrations -database "postgres://aash:@localhost:5432/implight?sslmode=disable" -verbose down
migrate_fix:
	migrate -path db/migrations -database "postgres://aash:@localhost:5432/implight?sslmode=disable" force $(VERSION)
