<center><img src="public/img/logo-text.png"></center>

Self-hostable web app for markdown note-taking, written in Golang and utilizing HTMX. Notes are stored in plain markdown files. Multiple user support via SQLite.

This project was heavily inspired by another excellent note taking application, [Flatnotes](https://github.com/Dullage/flatnotes).

![wip](/screenshots/wip.png)

![notes](/screenshots/notes-view.png)

![editor](/screenshots/editor-view.png)

![fullscreen](/screenshots/fullscreen-view.png)

## Development

Install [air](https://github.com/cosmtrek/air) for Golang live reload. Run the air server:

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
