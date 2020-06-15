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
