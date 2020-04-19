run:
	cd deployments/docker-compose && docker-compose up -d

restart:
	cd deployments/docker-compose && docker-compose restart

logs:
	cd deployments/docker-compose && docker-compose logs -f

build-api:
	docker build -t inflion-api -f build/api/Dockerfile .

gqlgen:
	go run github.com/99designs/gqlgen

gen:
	go generate ./...

sqlcgen:
	sqlc generate
