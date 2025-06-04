# Person List API

## Quick start
1. Run `git clone https://github.com/medvedevse/person-list-api.git`
2. Install dependencies `go build`
3. Set variables in .env (Examples of variables can be found in the .env.example file)
4. Run `go run cmd/app/main.go`

## Description
This is a simple API service with person data
Stack: Gin, Gorm, Postgres, Swaggo, Zap, godotenv

## Docs
Swagger: `http://localhost:8080/swagger/index.html`

## Endpoints
`GET /person` - get a full person list
`GET /person?age=29&gender=male&nationality=RU` - get a sorted person list
`GET /person?page=1&limit=3` - get a paginated person list
`POST /person` - add a new person
`PUT /person/:id` - update an existing person
`DELETE /person/:id` - delete an existing person
