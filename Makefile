prepare:
	docker-compose up -d --build

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

seed:
	docker-compose exec app go run seed/seed.go
