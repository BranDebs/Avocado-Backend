dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build

prod:
	docker-compose -f docker-compose.yml up --build

clean:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml down

.PHONY: dev prod clean
