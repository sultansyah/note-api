#
run database migration for mysql:
migrate -database "mysql:username:password@tcp(localhost:3306)/database_name" -path database/migrations up
