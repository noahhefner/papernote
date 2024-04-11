# Notes

!!! Work in progress, many features currently broken. !!!

Self-hostable, no-database web app for markdown note-taking, written in Golang and utilizing HTMX. Store notes in plain markdown files for easy import/export. Multiple user support.

![notes](/screenshots/notes-view.png)

![editor](/screenshots/editor.png)

## Development

Install [air](https://github.com/cosmtrek/air) for live reload. Run the air server:

```
air
```

Makefile commands:

| Command        | Action                                 |
| -------------- | -------------------------------------- |
| `make build`   | Create Docker container                |
| `make run`     | Run temporary Docker container         |
| `make clean`   | Delete Docker container                |
| `make publish` | Publish Docker container to Docker Hub |
