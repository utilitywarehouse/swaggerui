PHONY: build-ui
build-ui:
	cd app && npm run build

run:
	cd app && npm run build && cd .. && go run cmd/main.go