PHONY: build-ui
build-ui:
	cd app && npm run build

PHONY: run
run:
	cd app && npm run build && cd .. && go run cmd/main.go