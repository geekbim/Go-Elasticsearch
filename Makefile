.PHONY: docker-up
docker-up:
	docker-compose -f docker/docker-compose.yaml up

.PHONY: docker-down
docker-down:
	docker-compose -f docker/docker-compose.yaml down
	docker system prune --volume --force

.PHONY: up
up:
	go run -race cmd/elastic/main.go