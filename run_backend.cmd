@echo off
@echo setup environtment variable
SET fiber_db=host=localhost port=5432 user=postgres password=bismillah dbname=fiber sslmode=disable
cd backend
go run main.go