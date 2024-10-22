# Cursor Pagination Example

This repository contains an example implementation of cursor pagination in Go.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

  ```sh
  git clone https://github.com/eskimoburger/cursor-pagination.git
  cd cursor-pagination
  ```

1. Install dependencies:

  ```sh
  go mod tidy
  ```

### Running the Server

To start the server, run:

```sh
go run main.go
```

The server will start at `http://localhost:8000`.

### API Endpoints

- `GET /users`: Retrieves a paginated list of users.

### Example

To fetch the first page of users, you can use `curl`:

```sh
curl http://localhost:8000/users
```

### Example Request and Response

#### Request

```sh
curl http://localhost:8000/users?limit=10&cursor=eyJjcmVhdGVkX2F0IjoiMjAyMy0wMS0wMVQwMDowMDowMFoiLCJpZCI6MX0=
```

#### Response

```json
{
  "users": [
  {
    "id": 2,
    "name": "Bob",
    "created_at": "2023-01-02T00:00:00Z"
  },
  {
    "id": 3,
    "name": "Charlie",
    "created_at": "2023-01-03T00:00:00Z"
  }
  ],
  "next_cursor": "eyJjcmVhdGVkX2F0IjoiMjAyMy0wMS0wMlQwMDowMDowMFoiLCJpZCI6Mn0="
}
```

## Project Structure

- `controllers/`: Contains the HTTP handlers.
- `entities/`: Contains the data models.
- `repositories/`: Contains the data access logic.
- `services/`: Contains the business logic.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- Inspired by various cursor pagination implementations.
