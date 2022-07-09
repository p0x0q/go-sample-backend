serve:
	docker-compose up

generate:
	cd api && docker-compose up --build

command:
	docker-compose exec app /bin/sh
