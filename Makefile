build:
	docker-compose build --no-cache

run:
	docker-compose up --build
# docker-compose up --force-recreate --remove-orphans

stop:
	docker-compose down --remove-orphans

clear:
	docker-compose rm -f
	docker-compose pull