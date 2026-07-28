# CRUD User Golang

REST API in Go for user registration, lookup, update, deletion, and login. The project uses Gin to expose HTTP endpoints, MongoDB as the database, and JWT to generate an authentication token during login.

## How it works

1. Loads environment variables from the `.env` file.
2. Opens a MongoDB connection using the configured URL and database name.
3. Initializes the repository, service, and controller layers.
4. Registers the HTTP routes in Gin and starts the server on port `8080`.
5. Creates users after validating `email`, `password`, `name`, and `age`.
6. Persists users in MongoDB, storing the password as an MD5 hash.
7. Allows finding users by ID or e-mail, updating name/age, and deleting users.
8. During login, validates e-mail and password, generates a JWT valid for 24 hours, and returns the token in the `Authorization` header.

## Project structure

```text
crud-user-golang/
  main.go                                      # application entrypoint
  init_dependencies.go                         # main dependency initialization
  go.mod                                       # Go module and dependencies
  src/
    configuration/
      database/mongodb/                        # MongoDB connection
      logger/                                  # logger with zap
      validation/                              # Gin/validator validation and error translation
    controller/
      routes/                                  # HTTP route registration
      model/request/                           # API input contracts
      model/response/                          # API output contracts
      rest_err/                                # standardized error model
      createUser.go                            # create handler
      findUser.go                              # lookup handlers
      updateUser.go                            # update handler
      deleteUser.go                            # delete handler
      login_user.go                            # login handler
    model/
      service/                                 # business rules
      repository/                              # MongoDB access
      user_domain*.go                          # domain, password, and JWT token logic
    view/
      convert_domain_to_response.go            # domain-to-response conversion
```

## Requirements

- [Go](https://go.dev/dl/) 1.25+
- Local or remote MongoDB instance

## Environment variables

Create a `.env` file in the project root:

```env
MONGODB_URL=mongodb://localhost:27017
MONGODB_USER_DB=crud_user
MONGODB_USER_COLLECTION=users
JWT_SECRET_KEY=change-me
```

The `.env` file is not versioned.

## Running locally

Install dependencies:

```bash
go mod download
```

Start the application:

```bash
go run .
```

The API will be available at:

```text
http://localhost:8080
```

## API endpoints

| Method | Route | Description |
| --- | --- | --- |
| `POST` | `/createUser` | Creates a user |
| `POST` | `/login` | Authenticates a user and returns a JWT in the `Authorization` header |
| `GET` | `/getUserById/:userId` | Finds a user by MongoDB ID |
| `GET` | `/getUserByEmail/:userEmail` | Finds a user by e-mail using a JWT token |
| `PUT` | `/updateUser/:userId` | Updates name and/or age |
| `DELETE` | `/deleteUser/:userId` | Deletes a user |

## Request examples

### Create user

```bash
curl -X POST http://localhost:8080/createUser \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret@123",
    "name": "User Example",
    "age": 25
  }'
```

Response:

```json
{
  "id": "65f1c2a4b5c6d7e8f9012345",
  "email": "user@example.com",
  "name": "User Example",
  "age": 25
}
```

### Login

```bash
curl -i -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret@123"
  }'
```

The JWT token is returned in the `Authorization` header.

### Find user by ID

```bash
curl http://localhost:8080/getUserById/65f1c2a4b5c6d7e8f9012345
```

### Find user by e-mail

```bash
curl http://localhost:8080/getUserByEmail/user@example.com \
  -H "Authorization: Bearer <jwt-token>"
```

### Update user

```bash
curl -X PUT http://localhost:8080/updateUser/65f1c2a4b5c6d7e8f9012345 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated User",
    "age": 26
  }'
```

### Delete user

```bash
curl -X DELETE http://localhost:8080/deleteUser/65f1c2a4b5c6d7e8f9012345
```

## Validation rules

- `email`: required and must be a valid e-mail address.
- `password`: required, at least 6 characters long, and must contain at least one of these special characters: `!@#$%*`.
- `name`: required when creating a user, between 4 and 100 characters.
- `age`: required when creating a user, between 2 and 120.
- On update, `name` and `age` are optional, but they follow the same rules when provided.

## Notes

- The server uses a fixed `8080` port in the code.
- The `/getUserByEmail/:userEmail` route requires the `Authorization` header with a valid JWT.
- The other routes do not have authentication middleware in this version.
- There are no versioned tests in the project currently.
- For production, prefer a password hashing algorithm such as bcrypt or argon2 instead of MD5.
