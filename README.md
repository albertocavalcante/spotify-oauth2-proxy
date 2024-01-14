# Terraform Playlists

This project is a web application that serves OAuth2 authenticated requests. 
It's written in Go and uses Envoy Proxy OAuth2 filter for authentication handling.

## Prerequisites

- Go
- Docker

## Usage

To run the project, use the following command:

### Envoy Proxy

```sh
docker compose up
```

### Backend

```sh
go run ./...
```

Visit http://localhost:8000

## Project Structure

The project is structured as follows:

- `src/`: Contains the Go source files.
  - `oauth2/`: Contains the OAuth2 server implementation.
- `static/`: Contains static files served by the application.
  - `assets/`: Contains asset files.
  - `styles/`: Contains CSS stylesheets.

## License

Distributed under the MIT License. See `LICENSE` for more information.
