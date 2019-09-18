# komfy/api

[![Codacy](https://img.shields.io/codacy/grade/a4485f4982d54841930f0812b92f7c04.svg?style=flat-square)](https://app.codacy.com/project/komfy/api/dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/komfy/api)](https://goreportcard.com/report/github.com/komfy/api)
![GitHub last commit](https://img.shields.io/github/last-commit/komfy/api.svg?style=flat-square)

Komfy API built with Go, GraphQL, PostgreSQL and GORM.

## Useful links

[GraphQL API Prototype](https://app.graphqleditor.com/komfy/komfy-api)
[API Homepage](https://api.komfy.now.sh/)

## Local setup

1. Clone repository

```sh
git clone https://github.com/komfy/api.git
cd api
```

2. Install go modules

```sh
go mod download
```

3. Setup `.env`:

```
database=POSTGRES_DATABASE_URL
user_email=EMAIL_THAT_WILL_SEND_CONFIRMATION_EMAILS
pass_email=PASSWORD_FROM_EMAIL
secret=JWT_SECRET
```

Or use [`now secrets`](https://zeit.co/docs/v2/environment-variables-and-secrets)

3. Install Taskfile and Air:

```sh
go get -u -v github.com/go-task/task/cmd/task
curl -fLo ~/.local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
chmod +x ~/.local/bin/air
```

Be sure that `~/.local/bin/air` is in your path.

4. Run dev server

```
task dev
```

## Contributing

Coming soon...
