# User Authentication and Management API

This is a server-side boilerplate using Golang, echo framework, and mongoDB.

## Features

- **User Creation:** Allows the creation of new users with unique usernames.
- **User Login:** Authenticates users based on their username and password.
- **Token-based Authentication:** Utilizes JWT (JSON Web Token) for authentication of protected endpoints.

### Endpoint Details

- `/ping`: Endpoint to test the API availability.
- `/user/:name`: GET endpoint to retrieve user details based on username.
- `/login`: POST endpoint for user authentication and token generation.
- `/admin/user`: POST endpoint to create a new user (protected by authentication).
- `/admin` Endpoints: Protected endpoints accessible only with a valid JWT token.

## Usage

### Installation

1. Clone the repository.
2. Install required dependencies: `go get github.com/labstack/echo/v4 github.com/dgrijalva/jwt-go`
3. Run the Application: `go run index.go` to start the server on port 4000.

### API Endpoints

- `GET /ping`: Tests if the server is running.
- `GET /user/:name`: Retrieves user details by providing the username.
- `POST /login`: Authenticates users by providing a JSON payload with username and password.
- `POST /admin/user`: Creates a new user with a unique username. This endpoint is protected and requires a valid JWT token in the Authorization header (Bearer token).

### Protected Endpoints

Endpoints under `/admin` are protected and require a valid JWT token in the Authorization header to access.


### Sample Usage

**User Creation:**
```bash
curl -X POST http://localhost:4000/admin/user \
     -H "Authorization: Bearer <YOUR_JWT_TOKEN>" \
     -H "Content-Type: application/json" \
     -d '{
         "username": "new_user",
         "email": "new_user@example.com",
         "password": "password123"
     }'
```

**User Login:**
```bash
curl -X POST http://localhost:4000/login \
     -H "Content-Type: application/json" \
     -d '{
         "username": "amin",
         "password": "123"
     }'
```

Feel free to modify and expand upon this application to suit your specific requirements.

For further information on available endpoints and functionality, refer to the code documentation and Echo framework documentation.
