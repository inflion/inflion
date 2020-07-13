dev:
	docker run \
		--net host \
		--rm \
		-it \
		-w /go/src/github.com/inflion/inflion \
		-v `pwd`:/go/src/github.com/inflion/inflion \
		-v `pwd`/../.inflion.yaml:/root/.inflion.yaml \
		golang:1.14.4 \
		bash

up:
	cd deployments/docker-compose && docker-compose up -d

down:
	cd deployments/docker-compose && docker-compose down

restart:
	cd deployments/docker-compose && docker-compose restart

migrate-up:
	cd deployments/docker-compose && docker-compose run migration up

logs:
	cd deployments/docker-compose && docker-compose logs -f

sqlcgen:
	sqlc generate
