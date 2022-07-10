serve:
	docker-compose up --build

generate:
	docker-compose -f docker/openapi-generator/docker-compose.yml up --build

command:
	docker-compose exec app /bin/sh
