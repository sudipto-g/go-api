# Simple Go API - User Info Service

This is a simple Go-based API that listens on `localhost:8000` and responds to GET requests. The API retrieves a user's birthdate based on their `userid` from a mock database. Access to the API is secured using an authorization token, which must be provided in the request header.

## Features
- Exposes a simple GET endpoint that returns user data (user ID and birthdate).
- Authenticated by an authorization token sent via request headers.
- Written in Go (Golang).

## API Endpoint

### `/account/dob/`

Returns user information (birthdate) for the authorized user.

#### Request:

```http
GET http://localhost:8000/account/dob/?username=<username>
Authorization: <auth_token>
```

#### Response:

```json
{  
    "Date": "<DOB>",  
    "Code": 200  
}  
```

#### HTTP Status Codes:
- `200 OK`: Successfully retrieved user data.  
- `400 BadRequest`: Invalid user or token.  

## Authorization

The API requires a valid authorization token in the `Authorization` header. The format of the token should be:

```
Authorization: <auth_token>
```

For testing purposes, the token should be one of the following:
- `123sudo`: Returns birthdate `08-12-2000` for user ID `sudipto`  
- `456harsha`: Returns birthdate `23-11-1997` for user ID `harshada`  

Any other token (or missing token) will result in a `400 BadRequest` response.

## Mock Database

A mock database call (a golang map) is used to simulate retrieving user data. The following users are stored in the mock database:

| Token         | User ID    | Birthdate   |
|---------------|------------|-------------|
| `123sudo`     | `sudipto`  | `08-12-2000`|
| `456harsha`   | `harshada` | `23-11-1997`|

## How to Run

### Prerequisites

- Go should be installed. You can download Go from [https://golang.org/dl/](https://golang.org/dl/).
- A terminal or command-line interface to run the server.

### Steps to Run the API Server

1. Clone the repository.

   ```bash
   git clone https://github.com/sudipto-g/go-api.git
   cd go-api
   ```

2. Build and run the application:

   ```bash
   go run cmd/api/main.go
   ```

3. The server will start and listen on `http://localhost:8000`.

4. You can test the API with a valid authorization token using `curl` or `Postman`:

   **Example using curl:**

   ```bash
   curl -H "Authorization:<authtoken>" 'http://localhost:8000/account/dob/?username=<username>'
   ```

   **Response:**

   ```json
   {
     "Date": "<date_of_birth>",
     "Code": 200,
   }
   ```

## Directory Structure

```
.
├── README.md
├── api
│   └── api.go
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
└── internal
    ├── handlers
    │   ├── api.go
    │   └── get_dob.go
    ├── middleware
    │   └── authorization.go
    └── tools
        ├── db.go
        └── mockdb.go

```

---


### Explanation of the contents:
1. **Features and API Endpoints**: Describes what the API does and outlines the specific `GET /account/dob` endpoint.
2. **Authorization**: Explains how to pass the authorization token for accessing the API.
3. **Mock Database**: Simulates a database for user data and lists sample tokens and associated data.
4. **How to Run**: Provides clear steps for setting up and running the API on a local machine.
5. **Directory Structure**: A simple structure for the Go project files.

