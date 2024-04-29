<center><img src="public/img/logo-text.png"></center>

Self-hostable web app for markdown note-taking, written in Golang and utilizing HTMX. Notes are stored in plain markdown files. Multiple user support via SQLite.

This project was heavily inspired by another excellent note taking application, [Flatnotes](https://github.com/Dullage/flatnotes). It is very much a work in progress and lacks many features such as image support and the ability to rename notes.

![wip](/screenshots/wip.png)

## Notes Homepage

Landing page for browsing your notes. Hovering over a note title in the list shows a preview to the right. Quick actions (delete, fullscreen, edit) also appear next to the note title.

![notes](/screenshots/notes-view.png)

## Editor

No frills markdown editor. Write on the left, preview on the right. Thats it.

![editor](/screenshots/editor-view.png)

## Fullscreen View

Larger view of the rendered note.

![fullscreen](/screenshots/fullscreen-view.png)

## Deployment

Papernotes can be run as a Docker container via `docker-compose`:

```yml
# docker-compose.yml
---
version: "3"
services:
  papernote:
    image: noahhefner/papernote:latest
    container_name: papernote
    volumes:
      - /path/to/data:/data
    ports:
      - 8080:80
    user: 1000:1000
    environment:
      JWT_SECRET: ${JWT_SECRET}
```

```yml
# .env
JWT_SECRET=SUPER_SECRET_SECRET
```

The `/data` directory will be structured as follows:

```
/data
├── db
│   └── users.db
└── notes
    ├── user
    │   ├── cron-notes.md
    │   ├── lorum-ipsum.md
    │   └── yellowstone.md
    └── user2
        ├── user2-note1.md
        └── user2-note2.md
```

## Development

Dependencies:

- [Golang](https://go.dev/)
- [air](https://github.com/cosmtrek/air) (for live reload in development)
- [Docker](https://www.docker.com/)

The `test-data` directory provides a sample database and user directory with a few notes for testing purposes. `.air.toml` sets the `DATA_DIR` environment variable to `./test-data` for convenience. Note and database changes will be reflected in this directory. `/test-data` is in the `.gitignore`, so changes made while testing will not be committed to the remote.

Test user and password:

- Username: `user`
- Password: `pass`

Run the air server and access the application at `localhost:8080`:

```
air
```

## Docker

Helper commands for working with Docker are provided via `Makefile`.

| Command        | Action                                 |
| -------------- | -------------------------------------- |
| `make build`   | Create Docker container                |
| `make run`     | Run temporary Docker container         |
| `make clean`   | Delete Docker container                |
| `make publish` | Publish Docker container to Docker Hub |
