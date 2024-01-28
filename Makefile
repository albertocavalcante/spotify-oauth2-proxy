run:
	go run ./...
.PHONY: run

envoy-up:
	docker compose up -d
.PHONY: envoy-up

tailwind-watch:
	tailwindcss -i static/styles/main.css -o static/styles/output.css --watch=always
.PHONY: tailwind-watch

# https://tailwindcss.com/docs/optimizing-for-production
tailwind-minify:
	tailwindcss -i static/styles/main.css -o static/styles/output.css --minify
.PHONY: tailwind-minify
