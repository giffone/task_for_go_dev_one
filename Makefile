run:
	docker-compose rm -f
	docker-compose pull
	docker-compose up --build
# docker-compose up --force-recreate --remove-orphans

stop:
	docker-compose down --remove-orphans