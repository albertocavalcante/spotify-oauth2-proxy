run:
	go run ./...
.PHONY: run

tailwind-watch:
	tailwindcss -i static/styles/input.css -o static/styles/output.css --watch 
.PHONY: tailwind-watch
