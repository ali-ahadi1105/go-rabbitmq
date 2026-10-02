up:
	docker compose up -d
down:
	docker compose down
log:
	docker compose logs -f
build:
	docker compose build 
restart:
	docker compose down && docker compose up --build