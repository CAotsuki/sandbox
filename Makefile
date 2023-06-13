# go
run:
	go run cmd/main.go

curl:
	curl localhost:8080

# docker
up:
	docker compose -f build/docker-compose.yaml up

up-d:
	docker compose -f build/docker-compose.yaml up -d

build:
	docker compose -f build/docker-compose.yaml build --no-cache

log:
	docker compose -f build/docker-compose.yaml logs

down:
	docker compose -f build/docker-compose.yaml down

destroy:
	docker compose -f build/docker-compose.yaml down --rmi all --volumes --remove-orphans

# db
psql:
	docker compose -f build/docker-compose.yaml exec -it db psql -U todo todo

db-up:
	docker compose -f build/docker-compose.yaml up db

# test
test:
	zsh script/test.sh
