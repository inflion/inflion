up:
	docker-compose up -d

restart:
	docker-compose restart

destroy: stop
	docker-compose rm

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
