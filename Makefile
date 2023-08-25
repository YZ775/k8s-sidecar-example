.PHONY: build
build:
	cd app && docker build -t app .
	cd sidecar && docker build -t sidecar .

load:
	kind load docker-image app
	kind load docker-image sidecar
