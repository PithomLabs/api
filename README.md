# Komfy API

[![Codacy](https://img.shields.io/codacy/grade/a4485f4982d54841930f0812b92f7c04.svg?style=flat-square)](https://app.codacy.com/project/komfy/api/dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/komfy/api)](https://goreportcard.com/report/github.com/komfy/api)
![GitHub last commit](https://img.shields.io/github/last-commit/komfy/api.svg?style=flat-square)

Komfy API repository.

> [Komfy](https://komfy.now.sh) is a protected social network without annoying trackers and context ads. It is currently in active development.

## Stack

* **PostgresSQL** database
* **GraphQL** API
* **GORM** ORM

## Useful links

- [GraphQL API Prototype](https://app.graphqleditor.com/komfy-api/komfy-api)
- [API Homepage](https://api.komfy.now.sh/)

## Local setup

1. Clone repository

```sh
git clone https://github.com/komfy/api.git
cd api
```

2. Setup [Task](https://taskfile.dev) as shown [here](https://taskfile.dev/#/installation) 

2. Install go modules

```sh
go mod download
```

3. Setup `.env` as shown in [.env.example](https://github.com/komfy/api/blob/master/.env.example) 

3. Run dev server

```
task -w dev
```

## Contributing

Coming soon...
