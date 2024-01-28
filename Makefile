run:
	go run ./...
.PHONY: run

envoy-up:
	docker compose up -d
.PHONY: envoy-up

tailwind-watch:
	tailwindcss -i static/styles/main.css -o static/styles/output.css --watch
.PHONY: tailwind-watch
