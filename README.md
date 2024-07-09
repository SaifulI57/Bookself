# README

This project is a RESTful API built using the Gin framework for managing a library system. It includes functionalities for creating and retrieving users, borrowers, publishers, books, and authors.
Below are the details on how to set up and use the API.

## Table of Contents

-   [README](#readme)
    -   [Table of Contents](#table-of-contents)
    -   [Installation](#installation)
    -   [Usage](#usage)
        -   [Endpoints](#endpoints)
            -   [User Endpoints](#user-endpoints)
            -   [Borrower Endpoints](#borrower-endpoints)
            -   [Publisher Endpoints](#publisher-endpoints)
            -   [Book Endpoints](#book-endpoints)
            -   [Author Endpoints](#author-endpoints)
        -   [Examples](#examples)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/your-repo/library-api.git
    cd library-api
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

## Usage

### Endpoints

#### User Endpoints

-   **Create User**

    -   **URL:** `/createUser`
    -   **Method:** `POST`
    -   **Query Parameters:** `name`, `firstname`, `lastname`, `username`, `email`, `password`
    -   **Response:**
        ```json
        {
            "message": "User created successfully",
            "user": {
                "ID": 1,
                "Name": "John",
                "FirstName": "Doe",
                "LastName": "Smith",
                "Username": "johnsmith",
                "Email": "john@example.com",
                "Password": "password"
            }
        }
        ```

-   **Get All Users**

    -   **URL:** `/getUsers`
    -   **Method:** `GET`
    -   **Response:** Returns a list of all users.

-   **Get User**

    -   **URL:** `/getUser`
    -   **Method:** `GET`
    -   **Query Parameter:** `id`
    -   **Response:** Returns a specific user by ID.

-   **Update User**
    -   **URL:** `/updateUser`
    -   **Method:** `PUT`
    -   **Query Parameters:** `id`, `name`, `firstname`, `lastname`, `username`, `email`, `password`
    -   **Response:**
        ```json
        {
            "message": "User updated successfully"
        }
        ```

#### Borrower Endpoints

-   **Create Borrower**

    -   **URL:** `/createBorrower`
    -   **Method:** `POST`
    -   **Query Parameters:** `userId`, `endDate`, `startDate`, `bookId`
    -   **Response:**
        ```json
        {
            "message": "Borrower created successfully",
            "borrower": {
                "ID": 1,
                "User": {
                    /* user data */
                },
                "EndDate": "2024-12-31",
                "StartDate": "2024-01-01",
                "Book": {
                    /* book data */
                }
            }
        }
        ```

-   **Get All Borrowers**

    -   **URL:** `/getBorrowers`
    -   **Method:** `GET`
    -   **Response:** Returns a list of all borrowers.

-   **Get Borrower**
    -   **URL:** `/getBorrower`
    -   **Method:** `GET`
    -   **Query Parameter:** `id`
    -   **Response:** Returns a specific borrower by ID.

#### Publisher Endpoints

-   **Create Publisher**

    -   **URL:** `/createPublisher`
    -   **Method:** `POST`
    -   **Query Parameters:** `name`, `address`, `bookId`
    -   **Response:**
        ```json
        {
            "message": "Publisher created successfully",
            "publisher": {
                "ID": 1,
                "Name": "O'Reilly",
                "Address": "1005 Gravenstein Highway North, Sebastopol",
                "Books": [
                    /* book data */
                ]
            }
        }
        ```

-   **Get All Publishers**

    -   **URL:** `/getPublishers`
    -   **Method:** `GET`
    -   **Response:** Returns a list of all publishers.

-   **Get Publisher**
    -   **URL:** `/getPublisher`
    -   **Method:** `GET`
    -   **Query Parameter:** `id`
    -   **Response:** Returns a specific publisher by ID.

#### Book Endpoints

-   **Create Book**

    -   **URL:** `/createBook`
    -   **Method:** `POST`
    -   **Query Parameters:** `name`, `description`, `publisherId`, `authorId`, `isborrowed`
    -   **Response:**
        ```json
        {
            "message": "Book created successfully"
        }
        ```

-   **Get All Books**

    -   **URL:** `/getBooks`
    -   **Method:** `GET`
    -   **Response:** Returns a list of all books.

-   **Get Book**
    -   **URL:** `/getBook`
    -   **Method:** `GET`
    -   **Query Parameter:** `id`
    -   **Response:** Returns a specific book by ID.

#### Author Endpoints

-   **Create Author**

    -   **URL:** `/createAuthor`
    -   **Method:** `POST`
    -   **Query Parameters:** `name`, `email`, `bookId`
    -   **Response:**
        ```json
        {
            "message": "Author created successfully",
            "author": {
                "ID": 1,
                "Name": "John Doe",
                "Email": "john.doe@example.com",
                "Books": [
                    /* book data */
                ]
            }
        }
        ```

-   **Get All Authors**

    -   **URL:** `/getAuthors`
    -   **Method:** `GET`
    -   **Response:** Returns a list of all authors.

-   **Get Author**
    -   **URL:** `/getAuthor`
    -   **Method:** `GET`
    -   **Query Parameter:** `id`
    -   **Response:** Returns a specific author by ID.

### Examples

1. **Create a User**

    ```sh
    curl -X POST "http://localhost:8080/createUser?name=John&firstname=Doe&lastname=Smith&username=johnsmith&email=john@example.com&password=password"
    ```

2. **Get All Users**

    ```sh
    curl -X GET "http://localhost:8080/getUsers"
    ```

3. **Create a Book**
    ```sh
    curl -X POST "http://localhost:8080/createBook?name=GoProgramming&description=A book on Go programming&publisherId=1&authorId=1&isborrowed=false"
    ```
