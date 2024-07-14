# README

## Project Overview

This project provides a RESTful API for managing a library system using the Gin web framework. The API supports CRUD operations for managing authors, books, publishers, borrowers, and users. Each
entity is managed with concurrency-safe operations using the `sync.Mutex` for thread-safe data access.

## Project Structure

-   **Models**: Defined in `belajar_api/controllers`
-   **Request Handlers**: Defined in `request` package

## Data Structures

The project uses the following data structures:

-   `AuthorDB`
-   `BookDB`
-   `PublisherDB`
-   `BorrowerDB`
-   `UserDB`

Each of these data structures is equipped with a mutex to ensure thread-safe operations.

## API Endpoints

### Create Operations

#### Create User

-   **Endpoint**: `POST /user`
-   **Description**: Creates a new user.
-   **Parameters**:
    -   `name`
    -   `firstname`
    -   `username`
    -   `lastname`
    -   `email`
    -   `password`

#### Create Borrower

-   **Endpoint**: `POST /borrower`
-   **Description**: Creates a new borrower.
-   **Parameters**:
    -   `userId`
    -   `endDate`
    -   `startDate`
    -   `bookId`

#### Create Publisher

-   **Endpoint**: `POST /publisher`
-   **Description**: Creates a new publisher.
-   **Parameters**:
    -   `name`
    -   `address`

#### Create Book

-   **Endpoint**: `POST /book`
-   **Description**: Creates a new book.
-   **Parameters**:
    -   `name`
    -   `description`
    -   `publisherId`
    -   `authorId`

#### Create Author

-   **Endpoint**: `POST /author`
-   **Description**: Creates a new author.
-   **Parameters**:
    -   `name`
    -   `email`

### Read Operations

#### Get Users

-   **Endpoint**: `GET /users`
-   **Description**: Retrieves a list of all users.

#### Get User

-   **Endpoint**: `GET /user/:id`
-   **Description**: Retrieves a user by ID.

#### Get Borrowers

-   **Endpoint**: `GET /borrowers`
-   **Description**: Retrieves a list of all borrowers.

#### Get Borrower

-   **Endpoint**: `GET /borrower/:id`
-   **Description**: Retrieves a borrower by ID.

#### Get Publishers

-   **Endpoint**: `GET /publishers`
-   **Description**: Retrieves a list of all publishers.

#### Get Publisher

-   **Endpoint**: `GET /publisher/:id`
-   **Description**: Retrieves a publisher by ID.

#### Get Books

-   **Endpoint**: `GET /books`
-   **Description**: Retrieves a list of all books.

#### Get Book

-   **Endpoint**: `GET /book/:id`
-   **Description**: Retrieves a book by ID.

#### Get Authors

-   **Endpoint**: `GET /authors`
-   **Description**: Retrieves a list of all authors.

#### Get Author

-   **Endpoint**: `GET /author/:id`
-   **Description**: Retrieves an author by ID.

### Update Operations

#### Update User

-   **Endpoint**: `PUT /user/:id`
-   **Description**: Updates a user by ID.
-   **Parameters**:
    -   `name`
    -   `firstname`
    -   `lastname`
    -   `username`
    -   `email`
    -   `password`

#### Update Borrower

-   **Endpoint**: `PUT /borrower/:id`
-   **Description**: Updates a borrower by ID.
-   **Parameters**:
    -   `userId`
    -   `endDate`
    -   `startDate`
    -   `bookId`

#### Update Publisher

-   **Endpoint**: `PUT /publisher/:id`
-   **Description**: Updates a publisher by ID.
-   **Parameters**:
    -   `name`
    -   `address`

#### Update Book

-   **Endpoint**: `PUT /book/:id`
-   **Description**: Updates a book by ID.
-   **Parameters**:
    -   `name`
    -   `description`
    -   `publisherId`
    -   `authorId`

#### Update Author

-   **Endpoint**: `PUT /author/:id`
-   **Description**: Updates an author by ID.
-   **Parameters**:
    -   `name`
    -   `email`

### Delete Operations

#### Delete User

-   **Endpoint**: `DELETE /user/:id`
-   **Description**: Deletes a user by ID.

#### Delete Borrower

-   **Endpoint**: `DELETE /borrower/:id`
-   **Description**: Deletes a borrower by ID.

#### Delete Publisher

-   **Endpoint**: `DELETE /publisher/:id`
-   **Description**: Deletes a publisher by ID.

#### Delete Book

-   **Endpoint**: `DELETE /book/:id`
-   **Description**: Deletes a book by ID.

#### Delete Author

-   **Endpoint**: `DELETE /author/:id`
-   **Description**: Deletes an author by ID.

## Usage

To use the API, you need to run the server and make HTTP requests to the endpoints listed above.

### Example Request

To create a new user, you can send a POST request to `/user` with the required parameters.

```sh
curl -X POST "http://localhost:8080/user?name=John&firstname=Doe&lastname=Doe&username=johndoe&email=johndoe@example.com&password=secret"
```

## Concurrency

The project uses `sync.Mutex` to ensure thread-safe access to the in-memory data structures. This prevents race conditions when multiple goroutines try to read or write data simultaneously.

## Dependencies

-   `github.com/gin-gonic/gin`: Gin web framework

## Setup

1. Clone the repository
2. Install the dependencies: `go mod tidy`
3. Run the server: `go run main.go`

## Author

-   Saiful
