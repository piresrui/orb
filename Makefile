run-orb:
	docker run -it --env-file=./src/conf/config.yaml --rm -p 8080:8080 orb

build:
	docker build -t orb .

run:
	docker compose up -d
