# Notes

Work in progress

Self-hostable web app for notes. Backend in Go, frontend in Vue.js, stores markdown notes as plain markdown files.

## Development

Makefile commands:

| Command | Action |
|---------|--------|
| `make build` | Create Docker container |
| `make run` | Run temporary Docker container |
| `make clean` | Delete Docker container |
| `make publish` | Publish Docker container to Docker Hub |

To run just the Go backend:

```sh
cd server
go run .
```

To run just the Vue.js frontend:

```sh
cd app
npm run dev
```