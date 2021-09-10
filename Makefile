run:
	docker-compose up -d

stop:
	docker-compose stop

build:
	docker-compose build gomerch
	docker-compose up --no-deps -d gomerch
