serve:
	docker-compose up --build

generate:
	cd api && docker-compose up --build

command:
	docker-compose exec app /bin/sh
