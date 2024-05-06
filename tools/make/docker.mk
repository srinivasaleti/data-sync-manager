.PHONY: docker-run
docker-run:
docker-run: 
	docker compose -f tools/docker/docker-compose.yaml up --build