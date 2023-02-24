run-orb:
	docker run -it --env-file=./conf/config.yaml --rm orb

build-orb:
	docker build -t orb .

run-with-mock:
	docker compose up -d

go-generate:
	cd src; go generate ./...