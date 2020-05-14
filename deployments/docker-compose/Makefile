up:
	docker-compose up -d db
	sleep 3
	./migrate.sh
	docker-compose up -d hasura hasurats
	sleep 35
	cd ../../api/hasura/main && hasura metadata apply --admin-secret inflion
	docker-compose up -d

restart:
	docker-compose restart

destroy: stop
	docker-compose rm -f
	docker volume prune -f

build:
	docker-compose build

stop:
	docker-compose stop

ps:
	docker-compose ps

tail:
	docker-compose logs -f

shell:
	docker-compose exec api bash

gen:
	docker-compose exec api make gen

gqlgen:
	docker-compose exec api make gqlgen

sqlcgen:
	docker-compose exec api make sqlcgen

wire:
	docker-compose exec api make wire
