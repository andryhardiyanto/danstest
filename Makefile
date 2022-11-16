## docker-compose: build and run the docker compose
docker-compose: 
	echo "Pulling latest version"
	@docker-compose pull
	echo "Stop and remove the containers"
	@docker-compose rm -f
	echo "running the docker-compose container..."
	@docker-compose up -d --build
