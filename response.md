# Response
## A. Required Information
### A.1. Requirement Completion Rate
- [x] List all pharmacies open at a specific time and on a day of the week if requested.
  - Implemented at `GET /api/v1/pharmacies/by/time` API.
- [x] List all masks sold by a given pharmacy, sorted by mask name or price.
  - Implemented at `GET /api/v1/soldmasks/by/pharmacy` API.
- [x] List all pharmacies with more or less than x mask products within a price range.
  - Implemented at `GET /api/v1/pharmacies/masks/count` API.
- [x] The top x users by total transaction amount of masks within a date range.
  - Implemented at `GET /api/v1/usertrans/by/time` API.
- [x] The total number of masks and dollar value of transactions within a date range.
  - Implemented at `GET /api/v1/masktrans/by/time` API.
- [x] Search for pharmacies or masks by name, ranked by relevance to the search term.
  - Implemented at `GET /api/v1/pharmacies` API and `GET /api/v1/masks` .
- [x] Process a user purchases a mask from a pharmacy, and handle all relevant data changes in an atomic transaction.
  - Implemented at `POST /api/v1/purchase/mask` API.
### A.2. API Document
> Please describe how to use the API in the API documentation. You can edit by any format (e.g., Markdown or OpenAPI) or free tools (e.g., [hackMD](https://hackmd.io/), [postman](https://www.postman.com/), [google docs](https://docs.google.com/document/u/0/), or  [swagger](https://swagger.io/specification/)).

Import [kdan.swagger.json](doc/kdan/kdan.swagger.json) json file to Postman.
Or you can 

### A.3. Import Data Commands
Please run these commands to migrate the data into the database.
> Prerequisite: please install golang:1.20.1 & [migration CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
```bash
# Migrate postgresql
$ make migration_up
# Ingest JSON data to postgresql
$ go run ./cmd/client/ingestion.go
```
## B. Bonus Information

>  If you completed the bonus requirements, please fill in your task below.
### B.1. Test Coverage Report
Not implemented
### B.2. Dockerized
Please check my [Dockerfile](Dockerfile) / [docker-compose.yml](docker-compose.yml).

On the local machine, please follow the commands below to build it.

```bash
# Run docker-compose.local.yml
$ make up

$ make migration_up
# Ingest JSON data to postgresql
$ go run ./cmd/client/ingestion.go
```

Or you can build local docker-compose.local.yml and run go app instead

```bash
# Run docker-compose.local.yml
$ make local

# Run go app
$ make run

$ make migration_up
# Ingest JSON data to postgresql
$ go run ./cmd/client/ingestion.go
```

### B.3. Demo Site Url

Not implemented

## C. Other Information

### C.1. ERD

Not implemented. But you can examine the sql [tables](internal/data/migrations/000001_init_tables.up.sql) and [queries](internal/data/queries/).

### C.2. Technical Document

Not implemented.
- --
