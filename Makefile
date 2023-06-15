# go
run:
	go run cmd/main.go

build:
	go build -o sandbox cmd/main.go

curl:
	curl localhost:8080

test:
	zsh script/test.sh

# docker
docker:
	docker compose up

down:
	docker compose down

destroy:
	docker compose down --rmi all --volumes --remove-orphans

# postgresql
psql:
	docker compose exec -it db psql -U todo todo

# proxy
proxy-up:
	zsh script/cloud_sql_proxy.sh
