.PHONY: all
all: build run

.PHONY: build
build:
	mkdir -p dist && go build -o dist/env-templater ./cmd/env-templater

.PHONY: run
run:
	KAFKA_ZOOKEEPER_CONNECT=localhost:32181 KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:29092 KAFKA_BROKER_ID=2 KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 ./dist/env-templater -e KAFKA_ -t test/template.tmpl -o test/server.properties