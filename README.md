# Notes

Self-hostable web app for markdown note-taking, written in Golang and utilizing HTMX. Store notes in plain markdown files for easy import/export. Multiple user support via SQLite.

![wip](/screenshots/wip.jpg)

![notes](/screenshots/notes-view.png)

![editor](/screenshots/editor-view.png)

![fullscreen](/screenshots/fullscreen-view.png)

## Development

Install [air](https://github.com/cosmtrek/air) for live reload. Run the air server:

```
air
```

Test user and password:

- Username: `user`
- Password: `pass`

Makefile commands:

| Command        | Action                                 |
| -------------- | -------------------------------------- |
| `make build`   | Create Docker container                |
| `make run`     | Run temporary Docker container         |
| `make clean`   | Delete Docker container                |
| `make publish` | Publish Docker container to Docker Hub |
