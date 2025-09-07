# Blogger (Go + Fiber + GORM + SQLite)

A simple blog application with web UI and JSON API, built with Go, Fiber, GORM (SQLite), JWT auth, and server-side templates.

## Features
- Web UI for posts (list, view, create, edit, delete)
- JSON API for posts (CRUD)
- Registration/Login via forms (sets HttpOnly JWT cookie)
- API auth via `Authorization: Bearer <token>` or cookie
- Swagger docs at `/swagger/`

## Tech Stack
- Go, Fiber v2, html/template
- GORM with SQLite (`blog.db`)
- JWT (github.com/golang-jwt/jwt/v5)

## Prerequisites
- Go 1.20+ (recommended 1.22+)
- SQLite (no daemon required; uses a local file `blog.db`)

## Getting Started

### Build and Run
```bash
# build
go build ./...

# run
go run ./cmd/server
# server: http://127.0.0.1:3000
```

### Using Makefile (optional)
```bash
# build the project
make build

# start server in background
make run

# follow last 100 lines of server logs
make logs

# run the end-to-end smoke test (auth + CRUD)
make smoke

# stop the server
make stop

# clean temp artifacts
make clean
```

## Auth Flow
- Register at `/register` (form POST)
- Login at `/login` (form POST)
  - Sets `token` cookie (HttpOnly JWT)
- Protected routes require either:
  - `token` cookie (web), or
  - `Authorization: Bearer <token>` header (API)

## API Endpoints
Base URL: `http://127.0.0.1:3000`

- GET `/api/posts` → list posts
- GET `/api/posts/:id` → get post by id
- POST `/api/posts` → create post (JSON body)
- PUT `/api/posts/:id` → update post (JSON body)
- DELETE `/api/posts/:id` → delete post

Auth for API: include header `Authorization: Bearer <token>`.
You can obtain a token by logging in via `/login` and reading the `token` cookie.

### Example cURL (quick start)
```bash
# login (after registering via /register)
curl -i -c /tmp/c.txt -b /tmp/c.txt -X POST \
  -d "username=YOUR_USER" -d "password=YOUR_PASS" \
  http://127.0.0.1:3000/login

# extract token from cookie file
TOKEN=$(awk '($6=="token"){print $7}' /tmp/c.txt | tail -n1)

# create a post
curl -i -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" \
  -d '{"title":"Hello","content":"World"}' \
  http://127.0.0.1:3000/api/posts
```

## Smoke Test
A helper script `smoke.sh` automates:
- register → login → access protected `/`
- API posts: create → list → get → update → delete

Run it directly or via Makefile:
```bash
./smoke.sh
# or
make smoke
```

## Project Structure
```
cmd/server/          # main entry
internal/config/     # database init
internal/handlers/   # web + api handlers
internal/middlewares/# JWT middleware
internal/models/     # GORM models
internal/routes/     # route registration
internal/utils/      # JWT utils
templates/           # server-side templates
static/              # css/assets
```

## Notes
- SQLite file: `blog.db` created in project root on first run.
- Swagger served under `/swagger/` when server is running.
- For production, move `SecretKey` to an environment variable.

## License
MIT
