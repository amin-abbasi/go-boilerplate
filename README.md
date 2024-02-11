# Go Boilerplate Project

This Go boilerplate project demonstrates a simple user authentication system built using the Echo framework, MongoDB, and Redis.

## Project Structure

The project is organized into the following main files and packages:

### Main Files

- **`main.go`**: Initializes the server and sets up routes and middleware.

### Packages

#### Handlers (`handlers`)

Contains HTTP handlers responsible for managing HTTP endpoints and requests.

- **`handlers.go`**: Includes handlers for:
  - `GET /ping`: Responds with "pong".
  - `POST /login`: Validates user credentials and generates a JWT token.
  - `GET /user/:name`: Retrieves user details based on the provided username.
  - `POST /admin/user`: Creates a new user (requires authentication).

#### Middlewares (`middlewares`)

Contains middleware functions for authentication and authorization.

- **`middlewares.go`**
  - `Auth`: Middleware function for validating JWT tokens.

#### Models (`models`)

Contains data models and operations related to the user database.

- **`models.go`**
  - Defines the `User` struct representing a user with username, email, and password fields.
  - Manages the MongoDB database connection and collections.
  - Defines the `SECRET_KEY` used for JWT token generation.

## Running the Project

To run the project:

1. Ensure you have Go installed.
2. Clone this repository.
3. Navigate to the project directory.
4. Run `go run main.go` to start the server on port `4000`.

To run the project using Air for live reloading:

1. Install Air by running `go get -u github.com/cosmtrek/air`.
2. Create a configuration file named `air.toml` in the project root directory with the following content:

```toml
# air.toml

root = "."
tmp_dir = "tmp"
build_cmd = "go build -o ./tmp/main ."
run_cmd = "./tmp/main"
```
3. Run `air` in the project root directory to start the server with live reloading support.

## Usage

- **`GET /ping`**: Tests if the server is running.
- **`POST /login`**: Authenticates users by providing a JSON payload with username and password.
- **`GET /user/:name`**: Retrieves user details by providing the username.
- **`POST /admin/user`**: Creates a new user with a unique username. This endpoint is protected and requires a valid JWT token in the Authorization header (Bearer token).

### Notes

This boilerplate project showcases basic user authentication functionality using JWT tokens. Ensure to enhance security and error handling as needed based on your application's requirements.
