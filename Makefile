.PHONY: all
all: build run

.PHONY: build
build:
	mkdir -p dist && go build -o dist/env-templater ./cmd/env-templater

.PHONY: run
run:
	./dist/env-templater -t test/template.tmpl -o test/server.properties --extra-arg AdvertisedListeners=CLIENT://kafka-deployment-0.kafka-headless.kafka.svc.cluster.local:9092